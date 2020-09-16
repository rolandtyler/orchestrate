package utils

import (
	"context"
	"encoding/json"
	"html"
	"net/http"
	"net/url"

	"github.com/containous/traefik/v2/pkg/log"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/ethclient"
)

type apiError struct {
	Message string `json:"message"`
}

func ToFilters(values url.Values) map[string]string {
	filters := make(map[string]string)
	for key := range values {
		k := html.EscapeString(key)
		v := html.EscapeString(values.Get(key))
		if k != "" && v != "" {
			filters[k] = v
		}
	}
	return filters
}

func HandleStoreError(rw http.ResponseWriter, err error) {
	switch {
	case errors.IsAlreadyExistsError(err):
		WriteError(rw, err.Error(), http.StatusConflict)
	case errors.IsNotFoundError(err):
		WriteError(rw, err.Error(), http.StatusNotFound)
	case errors.IsDataError(err):
		WriteError(rw, err.Error(), http.StatusBadRequest)
	case err != nil:
		WriteError(rw, err.Error(), http.StatusInternalServerError)
	}
}

func WriteError(rw http.ResponseWriter, msg string, code int) {
	data, _ := json.Marshal(apiError{Message: msg})
	http.Error(rw, string(data), code)
}

func GetChainTip(ctx context.Context, ec ethclient.ChainLedgerReader, uris []string) (head uint64, err error) {
	var header *types.Header
	for _, uri := range uris {
		header, err = ec.HeaderByNumber(ctx, uri, nil)
		if err == nil {
			return header.Number.Uint64(), nil
		}
		log.FromContext(ctx).WithError(err).Warnf("failed to fetch chain id for URL %s", uri)
	}
	return
}
