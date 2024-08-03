package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDoors(t *testing.T) {

	assert.Equal(t, []direction{up, down, left}, openDoors("hijkl", nil))
	assert.Equal(t, []direction{up, left, right}, openDoors("hijkl", []direction{down}))
	assert.Equal(t, []direction{}, openDoors("hijkl", []direction{down, right}))
	assert.Equal(t, []direction{right}, openDoors("hijkl", []direction{down, up}))

}

func TestDfs(t *testing.T) {
	p1 := func(s string) string {
		p1, _ := dfs(s)
		return p1
	}

	assert.Equal(t, "DDRRRD", p1("ihgpwlah"))
	assert.Equal(t, "DDUDRLRRUDRD", p1("kglvqrro"))
	assert.Equal(t, "DRURDRUDDLLDLUURRDULRLDUUDDDRR", p1("ulqzkmiv"))

}
