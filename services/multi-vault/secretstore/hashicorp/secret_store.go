package hashicorp

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
)

// SecretStore wraps a HashiCorp client an manage the unsealing
type SecretStore struct {
	Client       *Hashicorp
	Config       *Config
	KeyBuilder   *multitenancy.KeyBuilder
	tokenWatcher *renewTokenWatcher
}

// NewSecretStore construct a new HashiCorp vault given a configfile or nil
func NewSecretStore(config *Config, keyBuilder *multitenancy.KeyBuilder) (*SecretStore, error) {
	client, err := NewVaultClient(config)
	if err != nil {
		return nil, errors.InternalError("HashiCorp: Could not start vault: %v", err)
	}

	tokenWatcher, err := newRenewTokenWatcher(client.Client, config.TokenFilePath)
	if err != nil {
		return nil, errors.InternalError("HashiCorp: Could not read token: %v", err)
	}

	if err := tokenWatcher.reloadToken(); err != nil {
		return nil, err
	}

	store := &SecretStore{
		Client:       client,
		Config:       config,
		KeyBuilder:   keyBuilder,
		tokenWatcher: tokenWatcher,
	}

	return store, nil
}

// Start starts a loop that will renew the token automatically
func (store *SecretStore) Start(ctx context.Context) error {
	go func() {
		err := store.tokenWatcher.Run(ctx)
		if err != nil {
			log.WithError(err).Fatal("token watcher routine has exited with errors")
		}
		log.Warn("token watcher routine has exited gracefully")
	}()

	return nil
}

// Store writes in the vault
func (store *SecretStore) Store(ctx context.Context, rawKey, value string) error {
	key, err := store.KeyBuilder.BuildKey(ctx, rawKey)
	if err != nil {
		return errors.FromError(err).ExtendComponent(component)
	}
	storedValue, ok, err := store.Client.Logical.Read(key)
	if err != nil {
		return errors.ConnectionError(err.Error()).ExtendComponent(component)
	}

	if ok {
		if storedValue == value {
			return nil
		}
		return errors.AlreadyExistsError("HashiCorp: A different secret already exists for key: %v", key).ExtendComponent(component)
	}

	err = store.Client.Logical.Write(key, value)
	if err != nil {
		return errors.ConnectionError(err.Error()).ExtendComponent(component)
	}
	return nil
}

// Load reads in the vault
func (store *SecretStore) Load(ctx context.Context, rawKey string) (value string, ok bool, e error) {
	allowedTenantIDs := multitenancy.AllowedTenantsFromContext(ctx)

	for _, tenant := range allowedTenantIDs {
		key := store.KeyBuilder.BuildKeyWithTenant(tenant, rawKey)

		v, ok, err := store.Client.Logical.Read(key)
		if err != nil {
			e = errors.ConnectionError(err.Error()).ExtendComponent(component)
		} else if ok {
			return v, ok, nil
		}
	}

	return "", false, e
}

// Delete removes a path in the vault
func (store *SecretStore) Delete(ctx context.Context, rawKey string) error {
	key, err := store.KeyBuilder.BuildKey(ctx, rawKey)
	if err != nil {
		return errors.FromError(err).ExtendComponent(component)
	}
	err = store.Client.Logical.Delete(key)
	if err != nil {
		return errors.ConnectionError(err.Error()).ExtendComponent(component)
	}
	return nil
}

// List returns the list of all secrets stored in the vault
func (store *SecretStore) List() (keys []string, err error) {
	keys, err = store.Client.Logical.List("")
	if err != nil {
		return []string{}, errors.ConnectionError(err.Error()).ExtendComponent(component)
	}
	return keys, err
}
