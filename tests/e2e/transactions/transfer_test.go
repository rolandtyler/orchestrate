package transactions

import (
	"github.com/consensys/orchestrate/pkg/types/api"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *transactionsTestSuite) TestTransfer() {
	s.Run("should send a transfer transaction successfully", func() {
		resp, err := s.env.Client.SendTransferTransaction(s.env.Ctx, &api.TransferRequest{
			ChainName: "",
			Labels:    nil,
			Params: api.TransferParams{
				Value:           "",
				Gas:             "",
				GasPrice:        "",
				GasFeeCap:       "",
				GasTipCap:       "",
				AccessList:      nil,
				TransactionType: "",
				From:            "",
				To:              "",
				GasPricePolicy:  api.GasPriceParams{},
			},
		})
		require.NoError(s.T(), err)

		assert.Equal(s.T(), resp.Jobs[0].Status, entities.StatusStarted)
	})
}
