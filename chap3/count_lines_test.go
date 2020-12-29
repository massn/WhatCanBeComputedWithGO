package chap3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountLines(t *testing.T) {
	length := countLines("hoge\nfuga\npoyo")
	assert.Equal(t, "3", length)
}
