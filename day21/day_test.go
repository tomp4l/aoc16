package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 steps
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`

func TestExample(t *testing.T) {
	ins, err := parse(example)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	t.Run("test parse", func(t *testing.T) {
		assert.Equal(t, ins[0], &swapPosition{from: 4, to: 0})
		assert.Equal(t, ins[1], &swapLetter{from: 'd', to: 'b'})
		assert.Equal(t, ins[2], &reverse{from: 0, to: 4})
		assert.Equal(t, ins[3], &rotateDir{left: true, amount: 1})
		assert.Equal(t, ins[4], &move{from: 1, to: 4})
		assert.Equal(t, ins[5], &move{from: 3, to: 0})
		assert.Equal(t, ins[6], &rotatePositionOf{letter: 'b'})
		assert.Equal(t, ins[7], &rotatePositionOf{letter: 'd'})
	})

	t.Run("matches example", func(t *testing.T) {
		assert.Equal(t, "decab", runAll("abcde", ins))
		assert.Equal(t, "abcde", reverseAll("decab", ins))

	})
}

func TestReverse(t *testing.T) {
	i := &reverse{2, 6}
	assert.Equal(t, "dcbheafg", i.execute("dcfaehbg"))
	assert.Equal(t, "dcfaehbg", i.reverse("dcbheafg"))

}

func TestPosition(t *testing.T) {
	i := &rotatePositionOf{letter: 'd'}
	assert.Equal(t, "ecabd", i.reverse("decab"))
}
