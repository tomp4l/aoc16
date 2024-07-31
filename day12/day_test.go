package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
jnz c 2
jnz 1 10
dec c
dec a
jnz 1 -4`

func TestExample(t *testing.T) {
	t.Run("parsing", func(t *testing.T) {
		res, err := parse(example)
		if assert.NoError(t, err) && assert.Len(t, res, 11) {
			copy, ok := res[0].(*copy)
			if assert.True(t, ok) {
				assert.Equal(t, 41, copy.i)
				assert.Equal(t, reg("a"), copy.dest)
				assert.Equal(t, reg(""), copy.r)
			}

			inc, ok := res[1].(*increase)
			if assert.True(t, ok) {
				assert.Equal(t, reg("a"), inc.reg)
			}

			dec, ok := res[3].(*decrease)
			if assert.True(t, ok) {
				assert.Equal(t, reg("a"), dec.reg)
			}

			jump, ok := res[4].(*jump)
			if assert.True(t, ok) {
				assert.Equal(t, reg("a"), jump.reg)
				assert.Equal(t, 2, jump.amount)
			}
		}
	})
	t.Run("run", func(t *testing.T) {
		p1, p2, err := (Day{}).Run(example)
		assert.NoError(t, err)
		assert.Equal(t, "42", p1)
		assert.Equal(t, "41", p2)

	})
}
