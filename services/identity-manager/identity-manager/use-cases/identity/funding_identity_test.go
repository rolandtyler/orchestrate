// +build unit

package identity

import (
	"context"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	testutils3 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/testutils"
	mock2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/client/mock"
	mock3 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/client/mock"
)

var (
	faucetNotFoundErr = errors.NotFoundError("not found faucet candidate")
)

func TestFundingIdentity_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegisterClient := mock2.NewMockChainRegistryClient(ctrl)
	mockTxSchedulerClient := mock3.NewMockTransactionSchedulerClient(ctrl)

	usecase := NewFundingIdentityUseCase(mockRegisterClient, mockTxSchedulerClient)

	t.Run("should trigger funding identity successfully", func(t *testing.T) {
		idenEntity := testutils3.FakeIdentity()
		chain := testutils3.FakeChain()
		faucet := testutils3.FakeFaucet()
		chainName := "besu"

		mockRegisterClient.EXPECT().GetChainByName(ctx, chainName).Return(chain, nil)
		mockRegisterClient.EXPECT().GetFaucetCandidate(ctx, ethcommon.HexToAddress(idenEntity.Address), chain.UUID).Return(faucet, nil)
		mockTxSchedulerClient.EXPECT().SendTransferTransaction(ctx, gomock.Any()).Return(nil, nil)
		err := usecase.Execute(ctx, idenEntity, chainName)

		assert.NoError(t, err)
	})

	t.Run("should do nothing if there is not faucet candidates", func(t *testing.T) {
		idenEntity := testutils3.FakeIdentity()
		chain := testutils3.FakeChain()
		chainName := "besu"

		mockRegisterClient.EXPECT().GetChainByName(ctx, chainName).Return(chain, nil)
		mockRegisterClient.EXPECT().GetFaucetCandidate(ctx, ethcommon.HexToAddress(idenEntity.Address), chain.UUID).
			Return(nil, faucetNotFoundErr)
		err := usecase.Execute(ctx, idenEntity, chainName)

		assert.NoError(t, err)
	})

	t.Run("should fail with same error if search chain fails", func(t *testing.T) {
		expectedErr := errors.NotFoundError("error")
		idenEntity := testutils3.FakeIdentity()
		chainName := "besu"

		mockRegisterClient.EXPECT().GetChainByName(ctx, chainName).Return(nil, expectedErr)
		err := usecase.Execute(ctx, idenEntity, chainName)

		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(fundingIdentityComponent), err)
	})

	t.Run("should fail with same error if get faucet candidate fails", func(t *testing.T) {
		expectedErr := errors.ConnectionError("error")
		idenEntity := testutils3.FakeIdentity()
		chain := testutils3.FakeChain()
		chainName := "besu"

		mockRegisterClient.EXPECT().GetChainByName(ctx, chainName).Return(chain, nil)
		mockRegisterClient.EXPECT().GetFaucetCandidate(ctx, ethcommon.HexToAddress(idenEntity.Address), chain.UUID).Return(nil, expectedErr)
		err := usecase.Execute(ctx, idenEntity, chainName)

		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(fundingIdentityComponent), err)
	})

	t.Run("should fail with same error if send funding transaction fails", func(t *testing.T) {
		expectedErr := errors.NotFoundError("error")
		idenEntity := testutils3.FakeIdentity()
		chain := testutils3.FakeChain()
		faucet := testutils3.FakeFaucet()
		chainName := "besu"

		mockRegisterClient.EXPECT().GetChainByName(ctx, chainName).Return(chain, nil)
		mockRegisterClient.EXPECT().GetFaucetCandidate(ctx, ethcommon.HexToAddress(idenEntity.Address), chain.UUID).Return(faucet, nil)
		mockTxSchedulerClient.EXPECT().SendTransferTransaction(ctx, gomock.Any()).Return(nil, expectedErr)
		err := usecase.Execute(ctx, idenEntity, chainName)

		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(fundingIdentityComponent), err)
	})
}
