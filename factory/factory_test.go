package factory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	//code here
	// Assert
	s.Assert().Equal(1, 1)
}
