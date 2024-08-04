package day19

import (
	"strconv"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	i, err := strconv.Atoi(input)
	if err != nil {
		return
	}
	p1 = strconv.Itoa(newRingAcross(i, false).winner())
	p2 = strconv.Itoa(newRingAcross(i, true).winner())

	return
}

type ringAcross struct {
	size    int
	current int
	steal   int
	elves   []bool
}

func newRingAcross(size int, across bool) *ringAcross {
	ring := new(ringAcross)
	ring.size = size
	ring.current = 0
	if across {
		ring.steal = size / 2
	} else {
		ring.steal = 1
	}
	ring.elves = make([]bool, size)

	return ring
}

func (r *ringAcross) winner() int {
	initialSize := r.size
	for {

		r.elves[r.steal] = true
		r.size -= 1

		if r.size == 1 {
			return r.current + 1
		}
		i := (r.current + 1) % initialSize
		for r.elves[i] {
			i++
			i %= initialSize
		}
		r.current = i

		skip := 1
		if r.size%2 == 0 {
			skip = 2
		}
		i = r.steal
		for skip > 0 {
			i++
			i %= initialSize
			if !r.elves[i] && r.current != i {
				skip--
			}
		}
		r.steal = i
	}
}
