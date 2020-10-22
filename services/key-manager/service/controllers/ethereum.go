package controllers

import (
	"encoding/json"
	"math/big"
	"net/http"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/keymanager"

	"github.com/gorilla/mux"
	jsonutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/encoding/json"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http/httputil"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/keymanager/ethereum"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/key-manager/key-manager/use-cases/ethereum"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/key-manager/service/formatters"
)

const Path = "/ethereum/accounts"

type EthereumController struct {
	ucs ethereum.UseCases
}

func NewEthereumController(ucs ethereum.UseCases) *EthereumController {
	return &EthereumController{ucs: ucs}
}

// Add routes to router
func (c *EthereumController) Append(router *mux.Router) {
	router.Methods(http.MethodPost).Path(Path).HandlerFunc(c.createAccount)
	router.Methods(http.MethodPost).Path(Path + "/import").HandlerFunc(c.importAccount)
	router.Methods(http.MethodPost).Path(Path + "/{address}/sign").HandlerFunc(c.signPayload)
	router.Methods(http.MethodPost).Path(Path + "/{address}/sign-transaction").HandlerFunc(c.signTransaction)
	router.Methods(http.MethodPost).Path(Path + "/{address}/sign-eea-transaction").HandlerFunc(c.signEEA)
	router.Methods(http.MethodPost).Path(Path + "/{address}/sign-tessera-transaction").HandlerFunc(c.signTessera)
}

// @Summary Creates a new Ethereum Account
// @Description Creates a new private key, stores it in the Vault and generates a public key given a chosen elliptic curve
// @Accept json
// @Produce json
// @Param request body ethereum.CreateETHAccountRequest true "Ethereum account creation request"
// @Success 200 {object} ethereum.ETHAccountResponse "Created Ethereum account"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts [post]
func (c *EthereumController) createAccount(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	ethAccountRequest := &types.CreateETHAccountRequest{}
	err := jsonutils.UnmarshalBody(request.Body, ethAccountRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	accountResponse, err := c.ucs.CreateAccount().Execute(ctx, ethAccountRequest.Namespace, "")
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatETHAccountResponse(accountResponse))
}

// @Summary Imports an Ethereum Account
// @Description Imports a private key, stores it in the Vault and generates a public key given a chosen elliptic curve
// @Accept json
// @Produce json
// @Param request body ethereum.ImportETHAccountRequest true "Ethereum account import request"
// @Success 200 {object} ethereum.ETHAccountResponse "Imported Ethereum account"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 422 {object} httputil.ErrorResponse "Invalid private key"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts/import [post]
func (c *EthereumController) importAccount(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	ethAccountRequest := &types.ImportETHAccountRequest{}
	err := jsonutils.UnmarshalBody(request.Body, ethAccountRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	accountResponse, err := c.ucs.CreateAccount().Execute(ctx, ethAccountRequest.Namespace, ethAccountRequest.PrivateKey)
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatETHAccountResponse(accountResponse))
}

// @Summary Signs an arbitrary message using an existing Ethereum account
// @Description Signs an arbitrary message using ECDSA and the private key of an existing Ethereum account
// @Accept json
// @Produce text/plain
// @Param request body keymanager.PayloadRequest true "Payload to sign"
// @Success 200 {string} string "Signed payload"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 404 {object} httputil.ErrorResponse "Account not found"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts/{address}/sign [post]
func (c *EthereumController) signPayload(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	signRequest := &keymanager.PayloadRequest{}
	err := jsonutils.UnmarshalBody(request.Body, signRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	address := mux.Vars(request)["address"]
	signature, err := c.ucs.SignPayload().Execute(ctx, address, signRequest.Namespace, signRequest.Data)
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_, _ = rw.Write([]byte(signature))
}

// @Summary Signs an Ethereum transaction using an existing account
// @Description Signs an Ethereum transaction using ECDSA and the private key of an existing account
// @Accept json
// @Produce text/plain
// @Param request body ethereum.SignETHTransactionRequest true "Ethereum transaction to sign"
// @Success 200 {string} string "Signed payload"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 404 {object} httputil.ErrorResponse "Account not found"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts/{address}/sign-transaction [post]
func (c *EthereumController) signTransaction(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	signRequest := &types.SignETHTransactionRequest{}
	err := jsonutils.UnmarshalBody(request.Body, signRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	address := mux.Vars(request)["address"]
	chainID, _ := new(big.Int).SetString(signRequest.ChainID, 10)
	tx := formatters.FormatSignETHTransactionRequest(signRequest)
	signature, err := c.ucs.SignTransaction().Execute(ctx, address, signRequest.Namespace, chainID, tx)
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_, _ = rw.Write([]byte(signature))
}

// @Summary Signs an EEA transaction using an existing account
// @Description Signs an EEA transaction using ECDSA and the private key of an existing account
// @Accept json
// @Produce text/plain
// @Param request body ethereum.SignETHTransactionRequest true "EEA transaction to sign"
// @Success 200 {string} string "Signed payload"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 404 {object} httputil.ErrorResponse "Account not found"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts/{address}/sign-eea-transaction [post]
func (c *EthereumController) signEEA(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	signRequest := &types.SignETHTransactionRequest{}
	err := jsonutils.UnmarshalBody(request.Body, signRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	address := mux.Vars(request)["address"]
	chainID, _ := new(big.Int).SetString(signRequest.ChainID, 10)
	tx := formatters.FormatSignETHTransactionRequest(signRequest)
	signature, err := c.ucs.SignTransaction().Execute(ctx, address, signRequest.Namespace, chainID, tx)
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_, _ = rw.Write([]byte(signature))
}

// @Summary Signs a Tessera transaction using an existing account
// @Description Signs a Tessera transaction using ECDSA and the private key of an existing account
// @Accept json
// @Produce text/plain
// @Param request body ethereum.SignTesseraTransactionRequest true "Private Tessera transaction to sign"
// @Success 200 {string} string "Signed payload"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 404 {object} httputil.ErrorResponse "Account not found"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /ethereum/accounts/{address}/sign-tessera-transaction [post]
func (c *EthereumController) signTessera(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	signRequest := &types.SignTesseraTransactionRequest{}
	err := jsonutils.UnmarshalBody(request.Body, signRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	address := mux.Vars(request)["address"]
	tx := formatters.FormatSignTesseraTransactionRequest(signRequest)
	signature, err := c.ucs.SignTesseraTransaction().Execute(ctx, address, signRequest.Namespace, tx)
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_, _ = rw.Write([]byte(signature))
}
