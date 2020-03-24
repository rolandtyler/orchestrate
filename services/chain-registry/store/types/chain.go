package types

import (
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/multitenancy"

	"github.com/containous/traefik/v2/pkg/config/dynamic"
	"github.com/containous/traefik/v2/pkg/tls"
	genuuid "github.com/satori/go.uuid"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
)

const component = "chain-registry.store"

type Chain struct {
	tableName struct{} `pg:"chains"` // nolint:unused,structcheck // reason

	UUID                      string     `json:"uuid" pg:",pk"`
	Name                      string     `json:"name"`
	TenantID                  string     `json:"tenantID"`
	URLs                      []string   `json:"urls" pg:"urls,array"`
	ListenerDepth             *uint64    `json:"listenerDepth,omitempty"`
	ListenerCurrentBlock      *uint64    `json:"listenerCurrentBlock,string,omitempty"`
	ListenerStartingBlock     *uint64    `json:"listenerStartingBlock,string,omitempty"`
	ListenerBackOffDuration   *string    `json:"listenerBackOffDuration,omitempty"`
	ListenerExternalTxEnabled *bool      `json:"listenerExternalTxEnabled,omitempty"`
	CreatedAt                 *time.Time `json:"createdAt"`
	UpdatedAt                 *time.Time `json:"updatedAt,omitempty"`
}

func (c *Chain) IsValid() bool {
	return c.Name != "" && c.TenantID != "" && len(c.URLs) != 0 && c.ListenerBackOffDuration != nil && *c.ListenerBackOffDuration != ""
}

func (c *Chain) SetDefault() {
	if c.UUID == "" {
		c.UUID = genuuid.NewV4().String()
	}
	if c.TenantID == "" {
		c.TenantID = multitenancy.DefaultTenantIDName
	}
	if c.ListenerDepth == nil {
		depth := uint64(0)
		c.ListenerDepth = &depth
	}
	if c.ListenerStartingBlock == nil {
		startingBlock := uint64(0)
		c.ListenerStartingBlock = &startingBlock
	}
	if c.ListenerCurrentBlock == nil {
		c.ListenerCurrentBlock = c.ListenerStartingBlock
	}
	if c.ListenerBackOffDuration == nil {
		backOffDuration := "1s"
		c.ListenerBackOffDuration = &backOffDuration
	}
	if c.ListenerExternalTxEnabled == nil {
		externalTxEnabled := false
		c.ListenerExternalTxEnabled = &externalTxEnabled
	}
}

func NewConfig() *dynamic.Configuration {
	return &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers:     make(map[string]*dynamic.Router),
			Middlewares: make(map[string]*dynamic.Middleware),
			Services:    make(map[string]*dynamic.Service),
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:  make(map[string]*dynamic.TCPRouter),
			Services: make(map[string]*dynamic.TCPService),
		},
		TLS: &dynamic.TLSConfiguration{
			Stores:  make(map[string]tls.Store),
			Options: make(map[string]tls.Options),
		},
	}
}

func BuildConfiguration(chains []*Chain) (*dynamic.Configuration, error) {
	config := NewConfig()

	for _, chain := range chains {
		chainKey := strings.Join([]string{chain.TenantID, chain.Name}, "-")

		config.HTTP.Routers[chainKey] = &dynamic.Router{
			EntryPoints: []string{"orchestrate"},
			Priority:    math.MaxInt32,
			Service:     chainKey,
			Rule:        fmt.Sprintf("Path(`/%s`)", chain.UUID),
			Middlewares: []string{
				"orchestrate-auth",
				"strip-path@internal",
				"orchestrate-ratelimit",
			},
		}

		servers := make([]dynamic.Server, 0)
		for _, chainURL := range chain.URLs {
			u, err := url.Parse(chainURL)
			if err != nil {
				return nil, errors.FromError(err).ExtendComponent(component)
			}

			servers = append(servers, dynamic.Server{
				Scheme: u.Scheme,
				URL:    chainURL,
			})
		}

		config.HTTP.Services[chainKey] = &dynamic.Service{
			LoadBalancer: &dynamic.ServersLoadBalancer{
				PassHostHeader: func(v bool) *bool { return &v }(false),
				Servers:        servers,
			},
		}
	}

	return config, nil
}