package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

func TestExample(t *testing.T) {
	s, err := parse(example)
	if assert.NoError(t, err) {
		t.Run("it parses", func(t *testing.T) {
			expected := &state{
				elevator: 1,
				floor1: &floor{map[string]chipGenerator{
					"hydrogen": {chip: true},
					"lithium":  {chip: true},
				}},
				floor2: &floor{map[string]chipGenerator{
					"hydrogen": {generator: true},
				}},
				floor3: &floor{map[string]chipGenerator{
					"lithium": {generator: true},
				}},
				floor4: &floor{map[string]chipGenerator{}},
			}
			assert.Equal(t, expected, s)
		})

		t.Run("it serialises", func(t *testing.T) {
			assert.Equal(t, "E1F1HMLMF2HGF3LGF4", s.serialise())
		})

		t.Run("is valid", func(t *testing.T) {
			assert.True(t, s.isValid())
		})

		t.Run("solve", func(t *testing.T) {
			assert.Equal(t, 11, solve(s))
		})
	}
}
