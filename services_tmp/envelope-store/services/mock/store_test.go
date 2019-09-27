package mock

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/envelope-store.git/services/testutils"
)

type MockEnvelopeStoreTestSuite struct {
	testutils.EnvelopeStoreTestSuite
}

func (s *MockEnvelopeStoreTestSuite) SetupTest() {
	s.Store = NewEnvelopeStore()
}

func TestMock(t *testing.T) {
	s := new(MockEnvelopeStoreTestSuite)
	suite.Run(t, s)
}
