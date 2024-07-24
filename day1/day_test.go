package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistances(t *testing.T) {

	t.Run("single distances", func(t *testing.T) {
		assertDistance(t, 1, "R1")
		assertDistance(t, 3, "L3")
		assertDistance(t, 123, "L123")
	})

	t.Run("multiple", func(t *testing.T) {
		assertDistance(t, 3, "R2, L1")
		assertDistance(t, 3, "R1, L2")
	})

	t.Run("examples", func(t *testing.T) {
		assertDistance(t, 5, "R2, L3")
		assertDistance(t, 2, "R2, R2, R2")
		assertDistance(t, 12, "R5, L5, R5, R3")

	})
}

func TestTwiceVisited(t *testing.T) {
	input := strings.Split("R8, R4, R4, R8", ", ")
	visited, err := visitedTwice(input)

	assert.NoError(t, err)
	assert.Equal(t, 4, visited)
}

func assertDistance(t testing.TB, expected int, raw string) {
	t.Helper()
	input := strings.Split(raw, ", ")
	dis, err := distance(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, dis)
}
