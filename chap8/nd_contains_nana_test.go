package chap8

import (
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

type ContainsNANATestSuite struct {
	suite.Suite
	length int
}

func (suite *ContainsNANATestSuite) SetupTest() {
	suite.length = 10
}

func TestContainsNANATestSuite(t *testing.T) {
	suite.Run(t, new(ContainsNANATestSuite))
}

func (suite *ContainsNANATestSuite) TestGAGA() {
	result := ndContainsNANA("ATCAGAGA")
	suite.Assertions.True(result)
}

func (suite *ContainsNANATestSuite) TestGAGAInRandomGenome() {
	result := getResult("GAGA", suite.length)
	suite.Assertions.True(result)
}

func getResult(target string, length int) bool {
	input := insertNANAToGenome(target, length)
	return ndContainsNANA(input)
}

func insertNANAToGenome(nana string, length int) string {
	b := getRandomGenome(length)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	p := r.Intn(length)
	return b[0:p] + nana + b[p:length]
}

func getRandomGenome(length int) string {
	g := []string{"A", "T", "G", "C"}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	genome := ""
	for i := 0; i < length; i++ {
		genome += g[r.Intn(4)]
	}
	return genome
}
