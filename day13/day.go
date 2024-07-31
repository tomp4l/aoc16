package day13

import (
	"strconv"
)

type Day struct{}

const (
	targetX     = 31
	targetY     = 39
	targetSteps = 50
)

func (Day) Run(input string) (p1 string, p2 string, err error) {

	magic, err := strconv.Atoi(input)
	if err != nil {
		return
	}
	a, b := distance(targetX, targetY, magic, targetSteps)
	p1 = strconv.Itoa(a)
	p2 = strconv.Itoa(b)

	return
}

func countBits(i int) int {
	c := 0
	m := 1
	for m <= i {
		if i&m == m {
			c++
		}
		m = m << 1
	}

	return c
}

func isWall(x, y, m int) bool {
	z := x*x + 3*x + 2*x*y + y + y*y
	b := countBits(z + m)
	return b%2 == 1
}

func distance(x, y, magic, steps int) (int, int) {
	type coord struct {
		x int
		y int
	}

	current := []coord{{1, 1}}
	seen := make(map[coord]bool)
	seen[current[0]] = true
	i := 0

	r1 := 0
	r2 := 0

	for {
		next := make([]coord, 0)
		if i == steps {
			r2 = len(seen)
		}
		for _, c := range current {
			if c.x == x && c.y == y {
				r1 = i
			}

			ns := []coord{{c.x + 1, c.y}, {c.x - 1, c.y}, {c.x, c.y + 1}, {c.x, c.y - 1}}
			for _, n := range ns {
				if n.x >= 0 && n.y >= 0 && !seen[n] && !isWall(n.x, n.y, magic) {
					next = append(next, n)
					seen[n] = true
				}
			}
		}
		if r1 != 0 && r2 != 0 {
			return r1, r2
		}
		current = next
		i++
	}
}
