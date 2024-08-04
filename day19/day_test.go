package day19

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	testCases := []struct {
		expectedNext   int
		expectedAcross int
		size           int
	}{
		{3, 2, 5},
		{5, 3, 6},
		{7, 5, 7},
		{1, 7, 8},
		{3, 9, 9},
		{5, 1, 10},
		{9, 13, 20},
		{29, 3, 30},
		{73, 19, 100},
	}
	for _, c := range testCases {
		t.Run(fmt.Sprintf("case %+v", c), func(t *testing.T) {
			assert.Equal(t, c.expectedNext, newRingAcross(c.size, false).winner())
			assert.Equal(t, c.expectedAcross, newRingAcross(c.size, true).winner())
		})
	}
}
