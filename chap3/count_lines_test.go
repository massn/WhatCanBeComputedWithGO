package chap3

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	length := countLines("hoge\nfuga\npoyo")
	assert.Equal(t, "3", length)
}

func TestSelfReflection(t *testing.T) {
	bytes, err := ioutil.ReadFile("count_lines.go")
	if err != nil {
		panic(err)
	}
	length := countLines(string(bytes))
	assert.Equal(t, "12", length)
}
