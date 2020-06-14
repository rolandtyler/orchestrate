// +build unit

package controllers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/encoding/json"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types"
	testutils3 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/service/formatters"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/service/testutils"
	testutils2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases/chains"
	mocks2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases/chains/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases/transactions"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases/transactions/mocks"
)

type transactionsControllerTestSuite struct {
	suite.Suite
	controller            *TransactionsController
	router                *mux.Router
	sendContractTxUseCase *mocks.MockSendContractTxUseCase
	sendDeployTxUseCase   *mocks.MockSendDeployTxUseCase
	sendTxUseCase         *mocks.MockSendTxUseCase
	getChainByNameUseCase *mocks2.MockGetChainByNameUseCase
	ctx                   context.Context
	tenantID              string
	chain                 *types.Chain
}

func (s *transactionsControllerTestSuite) SendContractTransaction() transactions.SendContractTxUseCase {
	return s.sendContractTxUseCase
}

func (s *transactionsControllerTestSuite) SendDeployTransaction() transactions.SendDeployTxUseCase {
	return s.sendDeployTxUseCase
}

func (s *transactionsControllerTestSuite) SendTransaction() transactions.SendTxUseCase {
	return s.sendTxUseCase
}

func (s *transactionsControllerTestSuite) GetChainByName() chains.GetChainByNameUseCase {
	return s.getChainByNameUseCase
}

var _ transactions.UseCases = &transactionsControllerTestSuite{}

func TestTransactionsController(t *testing.T) {
	s := new(transactionsControllerTestSuite)
	suite.Run(t, s)
}

func (s *transactionsControllerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.sendContractTxUseCase = mocks.NewMockSendContractTxUseCase(ctrl)
	s.sendDeployTxUseCase = mocks.NewMockSendDeployTxUseCase(ctrl)
	s.sendTxUseCase = mocks.NewMockSendTxUseCase(ctrl)
	s.getChainByNameUseCase = mocks2.NewMockGetChainByNameUseCase(ctrl)
	s.tenantID = "tenantId"
	s.chain = testutils3.FakeChain()
	s.ctx = context.WithValue(context.Background(), multitenancy.TenantIDKey, s.tenantID)

	s.router = mux.NewRouter()
	s.controller = NewTransactionsController(s, s)
	s.controller.Append(s.router)
}

func (s *transactionsControllerTestSuite) TestTransactionsController_Send() {

	s.T().Run("should execute request successfully", func(t *testing.T) {
		rw := httptest.NewRecorder()

		txRequest := testutils.FakeSendTransactionRequest(s.chain.Name)
		requestBytes, err := json.Marshal(txRequest)
		if err != nil {
			return
		}
		txRequestEntity := formatters.FormatSendTxRequest(txRequest)

		httpRequest := httptest.
			NewRequest(http.MethodPost, "/transactions/send", bytes.NewReader(requestBytes)).
			WithContext(s.ctx)

		testutils2.FakeTxRequestEntity()
		txRequestEntityResp := testutils2.FakeTxRequestEntity()

		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)
		
		s.sendContractTxUseCase.EXPECT().
			Execute(gomock.Any(), txRequestEntity, s.chain.UUID, s.tenantID).
			Return(txRequestEntityResp, nil)

		s.router.ServeHTTP(rw, httpRequest)

		response := formatters.FormatTxResponse(txRequestEntityResp, s.chain.Name)
		expectedBody, _ := json.Marshal(response)
		assert.Equal(t, string(expectedBody)+"\n", rw.Body.String())
		assert.Equal(t, http.StatusAccepted, rw.Code)
	})
	
	// Sufficient test to check that the mapping to HTTP errors is working. All other status code tests are done in integration tests
	s.T().Run("should fail with 422 if use case fails with InvalidParameterError", func(t *testing.T) {
		txRequest := testutils.FakeSendTransactionRequest(s.chain.Name)
		requestBytes, _ := json.Marshal(txRequest)
		txRequestEntity := formatters.FormatSendTxRequest(txRequest)
	
		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions/send",
			bytes.NewReader(requestBytes)).
			WithContext(s.ctx)
		
		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)
	
		s.sendContractTxUseCase.EXPECT().
			Execute(gomock.Any(), txRequestEntity, s.chain.UUID, s.tenantID).
			Return(nil, errors.InvalidParameterError("error"))
	
		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusUnprocessableEntity, rw.Code)
	})
	
	s.T().Run("should fail with Bad request if invalid format", func(t *testing.T) {
		txRequest := testutils.FakeSendTransactionRequest(s.chain.Name)
		txRequest.IdempotencyKey = ""
		requestBytes, _ := json.Marshal(txRequest)
	
		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions/send",
			bytes.NewReader(requestBytes)).WithContext(s.ctx)
	
		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
	})
	
	s.T().Run("should fail with 400 if request fails with InvalidParameterError for private txs", func(t *testing.T) {
		rw := httptest.NewRecorder()
		txRequest := testutils.FakeSendTesseraRequest(s.chain.Name)
		txRequest.Params.PrivateFrom = ""
		requestBytes, _ := json.Marshal(txRequest)
	
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions/send",
			bytes.NewReader(requestBytes)).
			WithContext(s.ctx)
	
		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
	})
}

func (s *transactionsControllerTestSuite) TestTransactionsController_Deploy() {
	urlPath := "/transactions/deploy-contract"

	s.T().Run("should execute request successfully", func(t *testing.T) {
		rw := httptest.NewRecorder()

		txRequest := testutils.FakeDeployContractRequest(s.chain.Name)
		requestBytes, _ := json.Marshal(txRequest)
		txRequestEntity := formatters.FormatDeployContractRequest(txRequest)

		httpRequest := httptest.NewRequest(http.MethodPost, urlPath, bytes.NewReader(requestBytes)).WithContext(s.ctx)

		txRequestEntityResp := testutils2.FakeTxRequestEntity()

		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)

		s.sendDeployTxUseCase.EXPECT().
			Execute(gomock.Any(), txRequestEntity, s.chain.UUID, s.tenantID).
			Return(txRequestEntityResp, nil)

		s.router.ServeHTTP(rw, httpRequest)

		response := formatters.FormatTxResponse(txRequestEntityResp, s.chain.Name)
		expectedBody, _ := json.Marshal(response)
		assert.Equal(t, string(expectedBody)+"\n", rw.Body.String())
		assert.Equal(t, http.StatusAccepted, rw.Code)
	})

	// Sufficient test to check that the mapping to HTTP errors is working. All other status code tests are done in integration tests
	s.T().Run("should fail with 422 if use case fails with InvalidParameterError", func(t *testing.T) {
		txRequest := testutils.FakeDeployContractRequest(s.chain.Name)
		requestBytes, _ := json.Marshal(txRequest)
		txRequestEntity := formatters.FormatDeployContractRequest(txRequest)

		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, urlPath, bytes.NewReader(requestBytes)).WithContext(s.ctx)

		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)

		s.sendDeployTxUseCase.EXPECT().
			Execute(gomock.Any(), txRequestEntity, s.chain.UUID, s.tenantID).
			Return(nil, errors.InvalidParameterError("error"))

		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusUnprocessableEntity, rw.Code)
	})

	s.T().Run("should fail with Bad request if invalid format", func(t *testing.T) {
		txRequest := testutils.FakeDeployContractRequest(s.chain.Name)
		txRequest.IdempotencyKey = ""
		requestBytes, _ := json.Marshal(txRequest)

		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, urlPath, bytes.NewReader(requestBytes)).WithContext(s.ctx)

		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
	})

	s.T().Run("should fail with 400 if request fails with InvalidParameterError for private txs", func(t *testing.T) {
		rw := httptest.NewRecorder()
		txRequest := testutils.FakeDeployContractRequest(s.chain.Name)
		txRequest.Params.PrivateFrom = "PrivateFrom"
		requestBytes, _ := json.Marshal(txRequest)

		httpRequest := httptest.NewRequest(http.MethodPost, urlPath, bytes.NewReader(requestBytes)).WithContext(s.ctx)

		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
	})
}

func (s *transactionsControllerTestSuite) TestTransactionsController_SendRaw() {

	s.T().Run("should execute request successfully", func(t *testing.T) {
		rw := httptest.NewRecorder()

		txRequest := testutils.FakeSendRawTransactionRequest(s.chain.Name)
		requestBytes, err := json.Marshal(txRequest)
		if err != nil {
			return
		}
		// txRequestEntity := formatters.FormatSendRawRequest(txRequest)

		httpRequest := httptest.
			NewRequest(http.MethodPost, "/transactions/send-raw", bytes.NewReader(requestBytes)).
			WithContext(s.ctx)

		testutils2.FakeTxRequestEntity()
		txRequestEntityResp := testutils2.FakeTxRequestEntity()

		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)

		s.sendTxUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any(), "", s.chain.UUID, s.tenantID).
			Return(txRequestEntityResp, nil)

		s.router.ServeHTTP(rw, httpRequest)

		response := formatters.FormatTxResponse(txRequestEntityResp, s.chain.Name)
		expectedBody, _ := json.Marshal(response)
		assert.Equal(t, string(expectedBody)+"\n", rw.Body.String())
		assert.Equal(t, http.StatusAccepted, rw.Code)
	})

	// Sufficient test to check that the mapping to HTTP errors is working. All other status code tests are done in integration tests
	s.T().Run("should fail with 422 if use case fails with InvalidParameterError", func(t *testing.T) {
		txRequest := testutils.FakeSendRawTransactionRequest(s.chain.Name)
		requestBytes, _ := json.Marshal(txRequest)
		txRequestEntity := formatters.FormatSendRawRequest(txRequest)

		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions/send-raw",
			bytes.NewReader(requestBytes)).
			WithContext(s.ctx)

		s.getChainByNameUseCase.EXPECT().
			Execute(gomock.Any(), s.chain.Name, s.tenantID).
			Return(s.chain, nil)

		s.sendTxUseCase.EXPECT().
			Execute(gomock.Any(), txRequestEntity, "", s.chain.UUID, s.tenantID).
			Return(nil, errors.InvalidParameterError("error"))

		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusUnprocessableEntity, rw.Code)
	})

	s.T().Run("should fail with Bad request if invalid format", func(t *testing.T) {
		txRequest := testutils.FakeSendRawTransactionRequest(s.chain.Name)
		txRequest.IdempotencyKey = ""
		requestBytes, _ := json.Marshal(txRequest)

		rw := httptest.NewRecorder()
		httpRequest := httptest.NewRequest(http.MethodPost, "/transactions/send-raw",
			bytes.NewReader(requestBytes)).WithContext(s.ctx)

		s.router.ServeHTTP(rw, httpRequest)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
	})
}
