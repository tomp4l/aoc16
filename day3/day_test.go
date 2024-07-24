package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangleFromString(t *testing.T) {

	tri, err := triangleFromString("   50   94    8")
	assert.NoError(t, err)
	assert.Equal(t, triangle{50, 94, 8}, tri)
}

func TestValidTriangles(t *testing.T) {
	valid := triangle{50, 94, 60}
	assert.True(t, valid.isValid())
	invalid := triangle{50, 94, 6}
	assert.False(t, invalid.isValid())
}
