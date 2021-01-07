package chap5

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"testing"

	"github.com/stretchr/testify/suite"
)

type BinaryIncrementerTestSuite struct {
	suite.Suite
	b *binaryIncrementer
}

func (suite *BinaryIncrementerTestSuite) SetupTest() {
	suite.b = NewBinaryIncrementer()
}

func TestBinaryIncremeenterTestSuite(t *testing.T) {
	suite.Run(t, new(BinaryIncrementerTestSuite))
}

func (suite *BinaryIncrementerTestSuite) TestOverflow() {
	inputTape := "x111x"
	accept, outTape := Run(suite.b, inputTape)
	suite.Assertions.True(accept)
	suite.Assertions.Equal("1000x", outTape)
}

func (suite *BinaryIncrementerTestSuite) TestIncrement1000times() {
	for i := 0; i < 1000; i++ {
		testIncrement(suite)
		suite.b.Reset()
	}
}

func testIncrement(suite *BinaryIncrementerTestSuite) {
	randomBits, incrementedRandomBits := makeRandomBitsAndIncremented()
	inputTape := "x" + randomBits + "x"
	accept, outTape := Run(suite.b, inputTape)
	suite.Assertions.True(accept)
	suite.Assertions.Equal("x"+incrementedRandomBits+"x", outTape)
}

func makeRandomBitsAndIncremented() (string, string) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Int63()
	randomBits := strconv.FormatInt(n, 2)
	if isOverFlowingBits(randomBits) {
		return makeRandomBitsAndIncremented()
	}
	incrementedRandomBits := strconv.FormatInt(n+1, 2)
	return randomBits, incrementedRandomBits
}

func isOverFlowingBits(bits string) bool {
	return !strings.Contains(bits, "0")
}
