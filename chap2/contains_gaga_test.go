package chap2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYes(t *testing.T) {
	result := countaintsGAGA("CTGAGAC")
	assert.Equal(t, "yes", result)
}
