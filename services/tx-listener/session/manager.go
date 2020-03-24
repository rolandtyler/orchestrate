package session

import (
	"context"
	"sync"

	"github.com/containous/traefik/v2/pkg/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/dynamic"
	provider "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-listener/providers"
)

type cancelableSession struct {
	session Session
	cancel  func()
}

type Manager struct {
	// Protect sessions mapping
	mux      *sync.Mutex
	sessions map[string]*cancelableSession

	// Wait Group to keep track of running sessions
	wg *sync.WaitGroup

	// Envelope used to create listening sessions
	builder Builder

	// Configuration
	currentConfiguration *dynamic.Configuration
	msgInput             chan *dynamic.Message

	// Configuration provider
	provider provider.Provider

	commands chan *Command

	errors chan error
}

func NewManager(sessionBuilder Builder, prvdr provider.Provider) *Manager {
	return &Manager{
		mux:                  &sync.Mutex{},
		sessions:             make(map[string]*cancelableSession),
		wg:                   &sync.WaitGroup{},
		currentConfiguration: &dynamic.Configuration{},
		msgInput:             make(chan *dynamic.Message),
		builder:              sessionBuilder,
		provider:             prvdr,
		commands:             make(chan *Command),
		errors:               make(chan error),
	}
}

func (m *Manager) Errors() <-chan error {
	return m.errors
}

func (m *Manager) Start(ctx context.Context) {
	m.run(ctx)
}

func (m *Manager) run(ctx context.Context) {
	defer func() {
		m.wg.Wait()
		log.WithoutContext().Infof("TxListener stopped")
	}()
	common.InParallel(
		// Start provider and close input channel when Provider is done
		func() {
			m.listenProvider(ctx)
		},
		// Listen configuration
		func() { m.listenConfiguration() },
		// Listen commands
		func() { m.listenCommands(ctx) },
		func() {
			<-ctx.Done()
			log.WithoutContext().Infof("I have to go... Stopping TxListener gracefully")
		},
	)
}

func (m *Manager) listenProvider(ctx context.Context) {
	log.Infof("Starting provider %T", m.provider)
	err := m.provider.Run(ctx, m.msgInput)
	if err != nil {
		log.FromContext(ctx).WithError(err).Errorf("error while listening provider")
	}
	close(m.msgInput)
}

func (m *Manager) listenConfiguration() {
	for msg := range m.msgInput {
		commands := CompareConfiguration(m.currentConfiguration, msg.Configuration)
		if len(commands) > 0 {
			for _, command := range commands {
				m.commands <- command
			}
			m.currentConfiguration = msg.Configuration
		}
	}
	close(m.commands)
}

func (m *Manager) listenCommands(ctx context.Context) {
	for command := range m.commands {
		m.executeCommand(ctx, command)
	}
}

func (m *Manager) executeCommand(ctx context.Context, command *Command) {
	switch command.Type {
	case START:
		m.runSession(ctx, command.Chain)
	case STOP:
		m.stopSession(command.Chain)
	case UPDATE:
		m.stopSession(command.Chain)
		m.runSession(ctx, command.Chain)
	default:
		log.WithoutContext().WithFields(logrus.Fields{
			"type":           command.Type,
			"chain.uuid":     command.Chain.UUID,
			"chain.tenantId": command.Chain.TenantID,
			"chain.name":     command.Chain.Name,
		}).Errorf("Unknown command")
	}
}

func (m *Manager) runSession(ctx context.Context, chain *dynamic.Chain) {
	// Build session
	s, err := m.builder.NewSession(chain)
	if err != nil {
		log.FromContext(ctx).WithError(err).Errorf("error while creating new session")
		return
	}

	// Make session cancelable session so we can stop it later on
	cancelableCtx, cancel := context.WithCancel(ctx)
	sess := &cancelableSession{
		session: s,
		cancel:  cancel,
	}

	// Add session
	m.addSession(chain.UUID, sess)

	// Start goroutine to run session
	m.wg.Add(1)
	go func() {
		logger := log.WithoutContext().WithField("session.chain.uuid", chain.UUID)
		logger.Infof("Session starts")
		err := sess.session.Run(log.With(cancelableCtx, log.Str("session.chain.uuid", chain.UUID)))
		m.removeSession(chain.UUID)
		if err != nil {
			log.FromContext(ctx).WithError(err).Errorf("error while running session")
		}
		logger.Infof("Session stopped")
		m.wg.Done()

	}()
}

func (m *Manager) stopSession(chain *dynamic.Chain) {
	sess, ok := m.getSession(chain.UUID)
	if ok {
		log.WithoutContext().WithField("session.chain.uuid", chain.UUID).Infof("Stopping session")
		sess.cancel()
	}
}

func (m *Manager) addSession(key string, sess *cancelableSession) {
	m.mux.Lock()
	m.sessions[key] = sess
	m.mux.Unlock()
}

func (m *Manager) removeSession(key string) {
	m.mux.Lock()
	delete(m.sessions, key)
	m.mux.Unlock()
}

func (m *Manager) getSession(key string) (sess *cancelableSession, ok bool) {
	m.mux.Lock()
	sess, ok = m.sessions[key]
	m.mux.Unlock()
	return
}