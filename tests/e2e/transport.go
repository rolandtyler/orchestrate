package e2e

import (
	"crypto/tls"
	"fmt"
	"github.com/consensys/orchestrate/pkg/toolkit/app/auth/utils"
	"net/http"
)

type testHttpTransport struct {
	token  string
	apiKey string
}

func NewTestHttpTransport(token, apiKey string) http.RoundTripper {
	return &testHttpTransport{
		token:  token,
		apiKey: apiKey,
	}
}

func (t *testHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	defaultTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	switch {
	case t.apiKey != "":
		req.Header.Add(utils.APIKeyHeader, t.apiKey)
	case t.token != "":
		req.Header.Add(utils.AuthorizationHeader, fmt.Sprintf("Bearer %s", t.token))
	}

	return defaultTransport.RoundTrip(req)
}
