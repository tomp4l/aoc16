package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNode(t *testing.T) {
	n, err := parseNode("/dev/grid/node-x1-y8     91T   70T    21T   76%")
	assert.NoError(t, err)
	assert.Equal(t, &node{x: 1, y: 8, size: 91, used: 70, available: 21}, n)
}
