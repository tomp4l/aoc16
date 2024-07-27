package day9

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type example struct {
	in  string
	out string
}

func TestExamples(t *testing.T) {
	cases := []example{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("case: %+v", c), func(t *testing.T) {
			assert.Equal(t, c.out, decompress(c.in))
		})
	}
}

type exampleLength struct {
	in     string
	length int
}

func TestExamples2(t *testing.T) {
	cases := []exampleLength{
		{"ADVENT", 6},
		{"X(8x2)(3x3)ABCY", len("XABCABCABCABCABCABCY")},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("case: %+v", c), func(t *testing.T) {
			assert.Equal(t, c.length, decompress2(c.in))
		})
	}
}
