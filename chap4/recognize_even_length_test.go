package chap4

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEven(t *testing.T) {
	result := recognizeEvenLength("hogefuga")
	assert.Equal(t, "yes", result)
}

func TestUndefined(t *testing.T) {
	resultChan := make(chan string, 1)
	go func() {
		resultChan <- recognizeEvenLength("hogefuga.")
	}()
	select {
	case <-resultChan:
		assert.Fail(t, "Not decidable instance.")
	case <-time.After(100 * time.Millisecond):
	}
}
