package chap3

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const gene = "CTGAGAT"
const geneticStringPath = "../chap2/geneticString.txt"

type YesTestSuite struct {
	suite.Suite
}

func TestYesTestSuite(t *testing.T) {
	suite.Run(t, new(YesTestSuite))
}

func (suite *YesTestSuite) TestGENE() {
	result := yes(gene)
	assert.Equal(suite.T(), "yes", result)
}

func (suite *YesTestSuite) TestGeneticString() {
	bytes, err := ioutil.ReadFile(geneticStringPath)
	if err != nil {
		panic(err)
	}
	result := yes(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}
