package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtensions(t *testing.T) {
	extended := extend(parse("..^^."), 3)

	assert.Equal(t, parse("..^^."), extended[0])
	assert.True(t, extended[1][0])
	assert.False(t, extended[1][1])
	assert.Equal(t, parse(".^^^^"), extended[1])
	assert.Equal(t, parse("^^..^"), extended[2])

	assert.Equal(t, 38, countSafe(extend(parse(".^^.^.^^^^"), 10)))
}
