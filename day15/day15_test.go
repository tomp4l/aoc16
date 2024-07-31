package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`

func TestParses(t *testing.T) {
	discs, err := parse(example)
	assert.NoError(t, err)

	if assert.Len(t, discs, 2) {
		assert.Equal(t, 1, discs[0].id)
		assert.Equal(t, 5, discs[0].positions)
		assert.Equal(t, 4, discs[0].start)
	}

}

func TestExample(t *testing.T) {
	p1, p2, err := (Day{}).Run(example)
	assert.NoError(t, err)
	assert.Equal(t, "5", p1)
	assert.Equal(t, "85", p2)
}
