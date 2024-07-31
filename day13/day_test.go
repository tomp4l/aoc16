package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, 0, countBits(0))
	assert.Equal(t, 1, countBits(1))
	assert.Equal(t, 1, countBits(2))
	assert.Equal(t, 2, countBits(3))
	assert.Equal(t, 1, countBits(4))
}

func TestWalls(t *testing.T) {
	magic := 10

	assert.False(t, isWall(0, 0, magic))
	assert.True(t, isWall(1, 0, magic))
	assert.False(t, isWall(0, 1, magic))
	assert.False(t, isWall(1, 1, magic))
	assert.False(t, isWall(2, 0, magic))
	assert.True(t, isWall(3, 0, magic))
}

func TestDistance(t *testing.T) {
	a, b := distance(7, 4, 10, 2)
	assert.Equal(t, 11, a)
	assert.Equal(t, 5, b)
}
