package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicDigit(t *testing.T) {
	k := new(basicKeypad)

	k.x = 1
	k.y = 1
	assert.Equal(t, 5, k.digit())

	k.x = 0
	k.y = 0
	assert.Equal(t, 1, k.digit())

	k.x = 2
	k.y = 2
	assert.Equal(t, 9, k.digit())
}

func TestBasicMove(t *testing.T) {
	t.Run("up", func(t *testing.T) {
		k := new(basicKeypad)
		k.y = 2

		assert.Equal(t, 7, k.digit())
		k.move(Up)
		assert.Equal(t, 4, k.digit())
		k.move(Up)
		assert.Equal(t, 1, k.digit())
		k.move(Up)
		assert.Equal(t, 1, k.digit())
	})

	t.Run("down", func(t *testing.T) {
		k := new(basicKeypad)

		assert.Equal(t, 1, k.digit())
		k.move(Down)
		assert.Equal(t, 4, k.digit())
		k.move(Down)
		assert.Equal(t, 7, k.digit())
		k.move(Down)
		assert.Equal(t, 7, k.digit())
	})

	t.Run("right", func(t *testing.T) {
		k := new(basicKeypad)

		assert.Equal(t, 1, k.digit())
		k.move(Right)
		assert.Equal(t, 2, k.digit())
		k.move(Right)
		assert.Equal(t, 3, k.digit())
		k.move(Right)
		assert.Equal(t, 3, k.digit())
	})

	t.Run("left", func(t *testing.T) {
		k := new(basicKeypad)
		k.x = 2

		assert.Equal(t, 3, k.digit())
		k.move(Left)
		assert.Equal(t, 2, k.digit())
		k.move(Left)
		assert.Equal(t, 1, k.digit())
		k.move(Left)
		assert.Equal(t, 1, k.digit())
	})
}

type digitCase struct {
	x int
	y int
	d int
}

func TestAdvancedDigit(t *testing.T) {

	cases := []digitCase{
		{2, 0, 1},
		{1, 1, 2},
		{2, 1, 3},
		{3, 1, 4},
		{0, 2, 5},
		{1, 2, 6},
		{2, 2, 7},
		{3, 2, 8},
		{4, 2, 9},
		{1, 3, 10},
		{2, 3, 11},
		{3, 3, 12},
		{2, 4, 13},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Case %+v", c), func(t *testing.T) {
			k := &advancedKeypad{c.x, c.y}
			assert.Equal(t, c.d, k.digit())
		})
	}
}

func TestExamples(t *testing.T) {
	input := `ULL
RRDDD
LURDL
UUUUD`
	d := new(Day)

	p1, p2, err := d.Run(input)

	assert.NoError(t, err)
	assert.Equal(t, p1, "1985")
	assert.Equal(t, p2, "5DB3")
}
