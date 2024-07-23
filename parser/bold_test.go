package parser

import (
	"testing"
	"wikinow/utils"

	"github.com/stretchr/testify/suite"
)

type BoldTestSuite struct {
	suite.Suite
	Data []utils.TestPair
}

func (suite *BoldTestSuite) SetupTest() {
	suite.Data = utils.GetTestData("..", "testdata", "bold.md")
	if len(suite.Data) == 0 {
		suite.T().Fatal("No test data found, ensure the file path is correct")
	}
}

// test if ParseBold returns the expected values
func (suite *BoldTestSuite) TestParseBoldReturnsExpected() {
	for _, pair := range suite.Data {
		received := ParseBold(pair.Input)
		suite.Equal(pair.Expected, received)
	}
}

func TestMain(t *testing.T) {
	suite.Run(t, new(BoldTestSuite))
}
