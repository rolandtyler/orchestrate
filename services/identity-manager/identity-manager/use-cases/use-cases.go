package usecases

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
)

//go:generate mockgen -source=use-cases.go -destination=mocks/use-cases.go -package=mocks

type IdentityUseCases interface {
	CreateIdentity() CreateIdentityUseCase
	SearchIdentity() SearchIdentitiesUseCase
}

type CreateIdentityUseCase interface {
	Execute(ctx context.Context, identity *entities.Identity, tenantID string) (*entities.Identity, error)
}

type SearchIdentitiesUseCase interface {
	Execute(ctx context.Context, filters *entities.IdentityFilters, tenants []string) ([]*entities.Identity, error)
}
