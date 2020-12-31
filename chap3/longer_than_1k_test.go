package chap3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1k(t *testing.T) {
	result := longerThan1k(strings.Repeat("0123456789", 100))
	assert.Equal(t, "yes", result)
}

func Test999(t *testing.T) {
	result := longerThan1k(strings.Repeat("123456789", 111))
	assert.Equal(t, "no", result)
}
