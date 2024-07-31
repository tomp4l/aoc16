package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHd(t *testing.T) {
	t.Run("fills initial", func(t *testing.T) {
		hd := newHd("100", 3)
		assert.True(t, hd.bits[0])
		assert.False(t, hd.bits[1])
		assert.False(t, hd.bits[2])
	})

	t.Run("fills rest", func(t *testing.T) {
		assert.Equal(t, "001", newHd("0", 3).String())
		assert.Equal(t, "100", newHd("1", 3).String())
		assert.Equal(t, "11111000000", newHd("11111", len("11111000000")).String())
		assert.Equal(t, "111100001010", newHd("111100001010", len("111100001010")).String())
		assert.Equal(t, "0010011000", newHd("0", 10).String())
	})
}
