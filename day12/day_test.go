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
	t.Run("run", func(t *testing.T) {
		p1, p2, err := (Day{}).Run(example)
		assert.NoError(t, err)
		assert.Equal(t, "42", p1)
		assert.Equal(t, "41", p2)

	})
}
