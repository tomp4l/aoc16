package day15

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	discs, err := parse(input)
	if err != nil {
		return
	}
	p1 = strconv.Itoa(alignment(discs))

	last := discs[len(discs)-1]
	discs = append(discs, disc{id: last.id + 1, positions: 11, start: 0})

	p2 = strconv.Itoa(alignment(discs))

	return
}

type disc struct {
	id        int
	positions int
	start     int
}

func parse(input string) ([]disc, error) {

	discs := make([]disc, 0)
	for _, l := range strings.Split(input, "\n") {
		split := strings.Split(l, " ")

		if len(split) != 12 {
			return nil, fmt.Errorf("unexpected line length, got %d for %s", len(split), l)
		}

		var err error
		disc := disc{}

		disc.id, err = strconv.Atoi(split[1][1:])
		if err != nil {
			return nil, err
		}

		disc.positions, err = strconv.Atoi(split[3])
		if err != nil {
			return nil, err
		}

		disc.start, err = strconv.Atoi(split[11][:len(split[11])-1])
		if err != nil {
			return nil, err
		}

		discs = append(discs, disc)
	}

	return discs, nil
}

func alignment(discs []disc) int {
	t := 0
	for {
		aligned := true
		t2 := t
		for _, d := range discs {
			t2++
			if d.position(t2) != 0 {
				aligned = false
				break
			}
		}
		if aligned {
			return t
		}
		t++
	}
}

func (d *disc) position(time int) int {
	return (d.start + time) % d.positions
}
