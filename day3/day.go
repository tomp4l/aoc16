package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	ts, err := parse(input)

	if err != nil {
		return
	}

	p1 = strconv.Itoa(countValid(ts))
	p2 = strconv.Itoa(countValid(transpose(ts)))

	return
}

type triangle struct {
	a int
	b int
	c int
}

func triangleFromString(input string) (t triangle, err error) {
	nonEmpty := make([]int, 0, 3)
	for _, p := range strings.Split(input, " ") {
		if p != "" {
			i, e := strconv.Atoi(p)
			if e != nil {
				err = e
				return
			}
			nonEmpty = append(nonEmpty, i)
		}
	}

	if len(nonEmpty) != 3 {
		err = fmt.Errorf("expected 3 sides but got %d", 3)
	} else {
		t = triangle{nonEmpty[0], nonEmpty[1], nonEmpty[2]}
	}

	return
}

func parse(input string) ([]triangle, error) {
	triangles := make([]triangle, 0)
	for _, l := range strings.Split(input, "\n") {
		t, err := triangleFromString(l)
		if err != nil {
			return nil, err
		}
		triangles = append(triangles, t)
	}

	return triangles, nil
}

func transpose(triangles []triangle) []triangle {
	transposed := make([]triangle, 0)
	for i := 0; i < len(triangles); i += 3 {
		t1 := triangles[i]
		t2 := triangles[i+1]
		t3 := triangles[i+2]

		transposed = append(transposed,
			triangle{t1.a, t2.a, t3.a},
			triangle{t1.b, t2.b, t3.b},
			triangle{t1.c, t2.c, t3.c})
	}
	return transposed
}

func countValid(ts []triangle) int {
	valid := 0
	for _, t := range ts {
		if t.isValid() {
			valid++
		}
	}
	return valid
}

func (t *triangle) isValid() bool {
	sides := []int{t.a, t.b, t.c}
	slices.Sort(sides)

	return sides[0]+sides[1] > sides[2]
}
