package controllers

import (
	"encoding/json"
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
// @Produce json
// @Param request body keymanager.PayloadRequest true "Payload to sign"
// @Success 200 {object} ethereum.ETHSignedPayloadResponse "Signed payload"
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
