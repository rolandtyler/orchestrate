package transactions

import (
	"github.com/consensys/orchestrate/tests/e2e"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/suite"
)

type transactionsTestSuite struct {
	suite.Suite
	env *e2e.Environment
}

func TestTransactions(t *testing.T) {
	s := new(transactionsTestSuite)
	suite.Run(t, s)
}

func (s *transactionsTestSuite) SetupSuite() {
	env, err := e2e.NewEnvironment()
	require.NoError(s.T(), err)
	s.env = env

	err = s.env.SetupResources()
	require.NoError(s.T(), err)
}

func (s *transactionsTestSuite) TearDownSuite() {
	err := s.env.TeardownResources()
	require.NoError(s.T(), err)
}
