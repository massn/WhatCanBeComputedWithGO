package chap3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var kString string = strings.Repeat("0123456789", 100)

func Test1k(t *testing.T) {
	result := longerThan1k(kString)
	assert.Equal(t, "no", result)
}

func Test1k1(t *testing.T) {
	result := longerThan1k(kString + ".")
	assert.Equal(t, "yes", result)
}
