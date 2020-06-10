package faucets

import (
	"encoding/json"
	"net/http"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/models"

	"github.com/gorilla/mux"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/multitenancy"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/chain-registry/utils"
)

// @Summary Retrieves a list of all registered faucet
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Success 200
// @Failure 404
// @Failure 500
// @Router /faucets [get]
func (h *controller) GetFaucets(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	faucets, err := h.getFaucetsUC.Execute(
		request.Context(),
		multitenancy.AllowedTenantsFromContext(request.Context()),
		utils.ToFilters(request.URL.Query()),
	)

	if err != nil {
		utils.HandleStoreError(rw, err)
		return
	}

	if len(faucets) == 0 {
		faucets = []*models.Faucet{}
	}

	_ = json.NewEncoder(rw).Encode(faucets)
}

// @Summary Retrieves a faucet by ID
// @Produce json
// @Param uuid path string true "ID of the faucet"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /faucets/{uuid} [get]
func (h *controller) GetFaucet(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	faucet, err := h.getFaucetUC.Execute(
		request.Context(),
		mux.Vars(request)["uuid"],
		multitenancy.AllowedTenantsFromContext(request.Context()),
	)

	if err != nil {
		utils.HandleStoreError(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(faucet)
}
