package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingInstruction(t *testing.T) {
	t.Run("parse rect", func(t *testing.T) {
		r, err := parseInstruction("rect 3x2")
		assert.NoError(t, err)

		rect, ok := r.(*rect)
		if !assert.True(t, ok) {
			t.Fatal()
		}
		assert.Equal(t, 3, rect.a)
		assert.Equal(t, 2, rect.b)
	})

	t.Run("parse rotate row", func(t *testing.T) {
		r, err := parseInstruction("rotate row y=1 by 4")
		assert.NoError(t, err)

		rect, ok := r.(*rotateRow)
		if !assert.True(t, ok) {
			t.Fatal()
		}
		assert.Equal(t, 1, rect.y)
		assert.Equal(t, 4, rect.b)
	})

	t.Run("parse rotate column", func(t *testing.T) {
		r, err := parseInstruction("rotate column x=1 by 2")
		assert.NoError(t, err)

		rect, ok := r.(*rotateColumn)
		if !assert.True(t, ok) {
			t.Fatal()
		}
		assert.Equal(t, 1, rect.x)
		assert.Equal(t, 2, rect.b)
	})
}

func TestApplyingInstructions(t *testing.T) {
	t.Run("test rect", func(t *testing.T) {
		r := &rect{2, 3}
		s := newScreen()
		r.apply(s)

		assert.True(t, s.pixels[0][0])
		assert.True(t, s.pixels[0][1])
		assert.False(t, s.pixels[0][2])
		assert.True(t, s.pixels[1][0])
		assert.True(t, s.pixels[2][0])
		assert.False(t, s.pixels[3][0])
		assert.True(t, s.pixels[2][1])
	})

	t.Run("test rotate row", func(t *testing.T) {
		s := newScreen()
		s.pixels[0][1] = true
		s.pixels[0][screenWidth-1] = true
		r := rotateRow{0, 1}
		r.apply(s)

		assert.True(t, s.pixels[0][0])
		assert.False(t, s.pixels[0][1])
		assert.True(t, s.pixels[0][2])
		assert.False(t, s.pixels[0][3])
	})

	t.Run("test rotate column", func(t *testing.T) {
		s := newScreen()
		s.pixels[1][0] = true
		s.pixels[screenHeight-1][0] = true
		r := rotateColumn{0, 1}
		r.apply(s)

		assert.True(t, s.pixels[0][0])
		assert.False(t, s.pixels[1][0])
		assert.True(t, s.pixels[2][0])
		assert.False(t, s.pixels[3][0])
	})
}
