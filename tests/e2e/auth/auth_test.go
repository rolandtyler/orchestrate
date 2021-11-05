package auth

import (
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/tests/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type authTestSuite struct {
	suite.Suite
	env *e2e.Environment
}

func TestAuth(t *testing.T) {
	s := new(authTestSuite)
	suite.Run(t, s)
}

func (s *authTestSuite) SetupSuite() {
	env, err := e2e.NewEnvironment()
	require.NoError(s.T(), err)
	s.env = env

	err = s.env.SetupResources()
	require.NoError(s.T(), err)
}

func (s *authTestSuite) TearDownSuite() {
	err := s.env.TeardownResources()
	require.NoError(s.T(), err)
}

func (s *authTestSuite) TestNoAuth() {
	s.Run("should fail to authenticate with status Unauthorized (401) if no auth method is provided", func() {
		orchestrateClient := client.NewHTTPClient(http.DefaultClient, client.NewConfig(s.env.Config.ApiURL, nil))

		_, err := orchestrateClient.GetSchedules(s.env.Ctx) // Call to random endpoint just to check we can't access it
		assert.True(s.T(), errors.IsUnauthorizedError(err))
	})
}
