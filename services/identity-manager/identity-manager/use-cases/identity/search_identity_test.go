// +build unit

package identity

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/identity-manager/identity-manager/parsers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/identity-manager/store/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/identity-manager/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/identity-manager/store/models/testutils"
)

func TestSearchIdentity_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	identityAgent := mocks.NewMockIdentityAgent(ctrl)
	mockDB.EXPECT().Identity().Return(identityAgent).AnyTimes()

	usecase := NewSearchIdentitiesUseCase(mockDB)

	tenantID := "tenantID"
	tenants := []string{tenantID}

	t.Run("should execute use case successfully", func(t *testing.T) {
		iden := testutils.FakeIdentityModel()
		
		filter := &entities.IdentityFilters{
			Aliases: []string{"alias1"},
		}
		
		identityAgent.EXPECT().Search(ctx, filter, tenants).Return([]*models.Identity{iden}, nil)

		resp, err := usecase.Execute(ctx, filter, tenants)

		assert.NoError(t, err)
		assert.Equal(t, parsers.NewIdentityEntityFromModels(iden), resp[0])
	})
	
	t.Run("should fail with same error if search identities fails", func(t *testing.T) {
		expectedErr := errors.NotFoundError("error")
		
		filter := &entities.IdentityFilters{
			Aliases: []string{"alias1"},
		}
		
		identityAgent.EXPECT().Search(ctx, filter, tenants).Return(nil, expectedErr)

		_, err := usecase.Execute(ctx, filter, tenants)

		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(createIdentityComponent), err)
	})
}
