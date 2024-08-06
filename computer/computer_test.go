package computer

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

const toggleExample = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

func TestExample(t *testing.T) {
	t.Run("parsing", func(t *testing.T) {
		res, err := ParseInstructions(example)
		if assert.NoError(t, err) && assert.Len(t, res, 11) {
			copy, ok := res[0].(*cpy)
			if assert.True(t, ok) {
				assert.Equal(t, iValue(41), copy.source)
				assert.Equal(t, reg("a"), copy.dest)
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
				assert.Equal(t, reg("a"), jump.check)
				assert.Equal(t, iValue(2), jump.amount)
			}
		}
	})
	t.Run("run", func(t *testing.T) {
		res, err := ParseInstructions(example)
		if assert.NoError(t, err) {
			comp := NewComputer(res)
			comp.RunAll()
			assert.Equal(t, 42, comp.Reg("a"))
		}
	})
}

func TestToggleExample(t *testing.T) {
	t.Run("parsing", func(t *testing.T) {
		res, err := ParseInstructions(toggleExample)
		if assert.NoError(t, err) && assert.Len(t, res, 7) {
			copy, ok := res[0].(*cpy)
			if assert.True(t, ok) {
				assert.Equal(t, iValue(2), copy.source)
				assert.Equal(t, reg("a"), copy.dest)
			}

			inc, ok := res[1].(*toggle)
			if assert.True(t, ok) {
				assert.Equal(t, reg("a"), inc.reg)
			}
		}
	})
	t.Run("run", func(t *testing.T) {
		res, err := ParseInstructions(toggleExample)
		if assert.NoError(t, err) {
			comp := NewComputer(res)
			comp.RunAll()
			assert.Equal(t, 3, comp.Reg("a"))

			assert.Equal(t, &increase{reg("a")}, comp.ins[3])
			assert.Equal(t, &jump{check: iValue(1), amount: reg("a")}, comp.ins[4])
		}
	})
}

func BenchmarkMultiplications(b *testing.B) {
	example := `cpy 10 a
cpy 1000 d
cpy 1000 c
inc a
dec c
jnz c -2
dec d
jnz d -5`

	ins, err := ParseInstructions(example)
	assert.NoError(b, err)

	for i := 0; i < b.N; i++ {
		c := NewComputer(ins)
		c.RunAll()
		assert.Equal(b, 1000010, c.Reg("a"))
	}
}
