package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

func TestExample(t *testing.T) {
	bots, err := parse(example)

	assert.NoError(t, err)

	t.Run("values", func(t *testing.T) {
		assert.Equal(t, 3, bots.bots[1].pending)
		assert.Equal(t, 0, bots.bots[1].high)
		assert.Equal(t, 0, bots.bots[1].low)

		assert.Equal(t, 0, bots.bots[2].pending)
		assert.Equal(t, 5, bots.bots[2].high)
		assert.Equal(t, 2, bots.bots[2].low)
	})

	t.Run("instructions", func(t *testing.T) {
		assert.Equal(t, 1, bots.bots[2].instruction.low)
		assert.Equal(t, true, bots.bots[2].instruction.lowBot)
		assert.Equal(t, 0, bots.bots[2].instruction.high)
		assert.Equal(t, true, bots.bots[2].instruction.highBot)

		assert.Equal(t, 1, bots.bots[1].instruction.low)
		assert.Equal(t, false, bots.bots[1].instruction.lowBot)
		assert.Equal(t, 0, bots.bots[1].instruction.high)
		assert.Equal(t, true, bots.bots[1].instruction.highBot)

		assert.Equal(t, 2, bots.bots[0].instruction.low)
		assert.Equal(t, false, bots.bots[0].instruction.lowBot)
		assert.Equal(t, 0, bots.bots[0].instruction.high)
		assert.Equal(t, false, bots.bots[0].instruction.highBot)
	})

	t.Run("comparers", func(t *testing.T) {
		assert.Equal(t, 2, bots.findComparer(2, 5))
		assert.Equal(t, 1, bots.findComparer(2, 3))
	})

	t.Run("run output", func(t *testing.T) {
		bots.runEnd()
		assert.Equal(t, 2, bots.outputs[1])
		assert.Equal(t, 3, bots.outputs[2])
		assert.Equal(t, 5, bots.outputs[0])
	})
}
