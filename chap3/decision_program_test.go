package chap3

import (
	"io/ioutil"
	"testing"

	"github.com/massn/WhatCanBeComputedWithGo/chap2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const gene = "CTGAGAT"
const geneticStringPath = "../chap2/geneticString.txt"
const containsGAGAPath = "../chap2/contains_gaga.go"
const longerThan1kPath = "longer_than_1k.go"
const yesPath = "yes.go"

type ContainsGAGATestSuite struct {
	suite.Suite
}

func TestContainsGAGATestSuite(t *testing.T) {
	suite.Run(t, new(ContainsGAGATestSuite))
}

func (suite *ContainsGAGATestSuite) TestGENE() {
	result := chap2.ContainsGAGA(gene)
	assert.Equal(suite.T(), "yes", result)
}

func (suite *ContainsGAGATestSuite) TestGeneticString() {
	bytes, err := ioutil.ReadFile(geneticStringPath)
	if err != nil {
		panic(err)
	}
	result := chap2.ContainsGAGA(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}

func (suite *ContainsGAGATestSuite) TestLongerThan1k() {
	bytes, err := ioutil.ReadFile(longerThan1kPath)
	if err != nil {
		panic(err)
	}
	result := chap2.ContainsGAGA(string(bytes))
	assert.Equal(suite.T(), "no", result)
}

func (suite *ContainsGAGATestSuite) TestContainsGAGA() {
	bytes, err := ioutil.ReadFile(containsGAGAPath)
	if err != nil {
		panic(err)
	}
	result := chap2.ContainsGAGA(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}

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

func (suite *YesTestSuite) TestContainsGAGA() {
	bytes, err := ioutil.ReadFile(containsGAGAPath)
	if err != nil {
		panic(err)
	}
	result := yes(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}

func (suite *YesTestSuite) TestYes() {
	bytes, err := ioutil.ReadFile(yesPath)
	if err != nil {
		panic(err)
	}
	result := yes(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}

type LongerThan1kTestSuite struct {
	suite.Suite
}

func TestLongerThan1kTestSuite(t *testing.T) {
	suite.Run(t, new(LongerThan1kTestSuite))
}

func (suite *LongerThan1kTestSuite) TestGENE() {
	result := longerThan1k(gene)
	assert.Equal(suite.T(), "no", result)
}

func (suite *LongerThan1kTestSuite) TestGeneticString() {
	bytes, err := ioutil.ReadFile(geneticStringPath)
	if err != nil {
		panic(err)
	}
	result := longerThan1k(string(bytes))
	assert.Equal(suite.T(), "yes", result)
}

func (suite *LongerThan1kTestSuite) TestContainsGAGA() {
	bytes, err := ioutil.ReadFile(containsGAGAPath)
	if err != nil {
		panic(err)
	}
	result := longerThan1k(string(bytes))
	assert.Equal(suite.T(), "no", result)
}

func (suite *LongerThan1kTestSuite) TestLongerThan1k() {
	bytes, err := ioutil.ReadFile(longerThan1kPath)
	if err != nil {
		panic(err)
	}
	result := longerThan1k(string(bytes))
	assert.Equal(suite.T(), "no", result)
}
