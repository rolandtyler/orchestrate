package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/pkg/types/api"
	"github.com/consensys/orchestrate/tests/e2e"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
)

func (s *authTestSuite) TestJWT() {
	s.Run("should create a resource successfully with JWT", func() {
		jwtToken, err := s.getJWT("https://orchestrate.consensys.net")
		require.NoError(s.T(), err)

		orchestrateClient := client.NewHTTPClient(
			&http.Client{Transport: e2e.NewTestHttpTransport(jwtToken, "")},
			client.NewConfig(s.env.Config.ApiURL, nil),
		)

		resp, err := orchestrateClient.CreateSchedule(s.env.Ctx, &api.CreateScheduleRequest{})
		require.NoError(s.T(), err)

		assert.Equal(s.T(), fmt.Sprintf("%s@clients", s.env.Config.AuthJWTClientID), resp.TenantID)
	})

	s.Run("should fail to authenticate using invalid JWT with status Unauthorized (401)", func() {
		orchestrateClient := client.NewHTTPClient(
			&http.Client{Transport: e2e.NewTestHttpTransport("invalidToken", "")},
			client.NewConfig(s.env.Config.ApiURL, nil),
		)

		_, err := orchestrateClient.GetAccount(s.env.Ctx, "accountAddress")
		require.True(s.T(), errors.IsUnauthorizedError(err))
	})

	// TODO: Delete accounts, currently not implemented
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *authTestSuite) getJWT(audience string) (string, error) {
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(map[string]interface{}{
		"client_id":     s.env.Config.AuthJWTClientID,
		"client_secret": s.env.Config.AuthJWTClientSecret,
		"audience":      audience,
		"grant_type":    "client_credentials",
	})

	resp, err := http.DefaultClient.Post(s.env.Config.AuthJWTTokenURL, jsonrpc.ContentType, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	acessToken := &accessTokenResponse{}
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(acessToken); err != nil {
			return "", err
		}

		return acessToken.AccessToken, nil
	}

	// Read body
	respMsg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return "", fmt.Errorf(string(respMsg))
}
