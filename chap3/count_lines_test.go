package chap3

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CountLinesTestSuite struct {
	suite.Suite
}

func TestCountLinesTestSuite(t *testing.T) {
	suite.Run(t, new(CountLinesTestSuite))
}

func (suite *CountLinesTestSuite) TestNormal() {
	length := countLines("hoge\nfuga\npoyo")
	assert.Equal(suite.T(), "3", length)
}

func (suite *CountLinesTestSuite) TestSelfReflection() {
	bytes, err := ioutil.ReadFile("count_lines.go")
	if err != nil {
		panic(err)
	}
	length := countLines(string(bytes))
	assert.Equal(suite.T(), "12", length)
}
