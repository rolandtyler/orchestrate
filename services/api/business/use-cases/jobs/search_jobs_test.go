// +build unit

package jobs

import (
	"context"
	"github.com/gofrs/uuid"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/parsers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models/testutils"
)

func TestSearchJobs_Execute(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	mockJobDA := mocks.NewMockJobAgent(ctrl)

	mockDB.EXPECT().Job().Return(mockJobDA).AnyTimes()

	usecase := NewSearchJobsUseCase(mockDB)

	tenantID := "tenantID"

	t.Run("should execute use case successfully", func(t *testing.T) {
		txHash := common.HexToHash("0x1")
		jobs := []*models.Job{testutils.FakeJobModel(0)}
		chainUUID := uuid.Must(uuid.NewV4()).String()
		filters := &entities.JobFilters{
			TxHashes:  []string{txHash.String()},
			ChainUUID: chainUUID,
		}

		expectedResponse := []*entities.Job{parsers.NewJobEntityFromModels(jobs[0])}
		mockJobDA.EXPECT().Search(ctx, filters, []string{tenantID}).Return(jobs, nil)
		jobResponse, err := usecase.Execute(ctx, filters, []string{tenantID})

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, jobResponse)
	})

	t.Run("should fail with same error if search fails for jobs", func(t *testing.T) {
		filters := &entities.JobFilters{}
		expectedErr := errors.NotFoundError("error")

		mockJobDA.EXPECT().Search(ctx, filters, []string{tenantID}).Return(nil, expectedErr)

		response, err := usecase.Execute(ctx, filters, []string{tenantID})

		assert.Nil(t, response)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(searchJobsComponent), err)
	})

	t.Run("should fail with invalid parameter in case invalid txHashes", func(t *testing.T) {
		filters := &entities.JobFilters{
			TxHashes: []string{"axasad"},
		}

		_, err := usecase.Execute(ctx, filters, []string{tenantID})

		assert.True(t, errors.IsInvalidParameterError(err))
	})

	t.Run("should fail with invalid parameter in case invalid chainUUID", func(t *testing.T) {
		filters := &entities.JobFilters{
			ChainUUID: "axasad",
		}

		_, err := usecase.Execute(ctx, filters, []string{tenantID})

		assert.True(t, errors.IsInvalidParameterError(err))
	})
}