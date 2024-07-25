package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	example := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	p1, p2, err := (Day{}).Run(example)

	assert.NoError(t, err)
	assert.Equal(t, "easter", p1)
	assert.Equal(t, "advent", p2)

}
