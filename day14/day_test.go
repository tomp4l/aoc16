package day14

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	p1, p2, _ := (Day{}).Run("abc")
	assert.Equal(t, "22728", p1)
	assert.Equal(t, "22551", p2)
}

func TestStretched(t *testing.T) {
	assert.True(t, strings.HasPrefix(stretched("abc", 0), "a107ff"), stretched("abc", 0))
}
