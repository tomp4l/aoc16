package day24

import (
	"math"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	maze, err := parse(input)
	if err != nil {
		return
	}
	p1 = strconv.Itoa(maze.shortestPath(false))
	p2 = strconv.Itoa(maze.shortestPath(true))

	return
}

func parse(input string) (*maze, error) {
	m := new(maze)
	m.nums = make(map[int]coord)
	m.paths = make(map[coord]bool)

	x := 0
	y := 0
	for _, l := range strings.Split(input, "\n") {
		for _, c := range l {
			if c != '#' {
				coord := coord{x, y}
				m.paths[coord] = true
				n, err := strconv.Atoi(string(c))
				if err == nil {
					m.nums[n] = coord
					if n == 0 {
						m.start = coord
					}
					if n > m.maxNum {
						m.maxNum = n
					}
				}
			}
			x++
		}

		y++
		x = 0
	}

	distances := make(map[int]map[int]int)
	for i := 0; i <= m.maxNum; i++ {
		distances[i] = make(map[int]int)
	}

	for i := 0; i <= m.maxNum; i++ {

		for j := i + 1; j <= m.maxNum; j++ {
			distance := 0
			visited := make(map[coord]bool)
			start := m.nums[i]
			end := m.nums[j]
			candidates := []coord{start}
			visited[start] = true

			for {
				newCandidates := make([]coord, 0)
				for _, c := range candidates {
					if c == end {
						distances[i][j] = distance
						distances[j][i] = distance
						goto next
					}
					neighbours := []coord{
						{c.x + 1, c.y},
						{c.x - 1, c.y},
						{c.x, c.y + 1},
						{c.x, c.y - 1},
					}
					for _, n := range neighbours {
						if m.paths[n] {
							newCandidates = append(newCandidates, n)
						}
					}
				}

				candidates = make([]coord, 0)
				for _, n := range newCandidates {
					if !visited[n] {
						visited[n] = true
						candidates = append(candidates, n)
					}
				}
				distance++
			}
		next:
		}
	}
	m.distances = distances

	return m, nil
}

type coord struct {
	x int
	y int
}

type maze struct {
	paths     map[coord]bool
	nums      map[int]coord
	maxNum    int
	start     coord
	distances map[int]map[int]int
}

func (m *maze) shortestPath(retZero bool) int {

	type state struct {
		visited  map[int]bool
		distance int
		current  int
	}

	startingState := state{make(map[int]bool), 0, 0}
	startingState.visited[0] = true
	states := []state{startingState}
	minDistance := math.MaxInt

	for len(states) > 0 {
		newStates := make([]state, 0)
		for _, s := range states {
			if len(s.visited) == m.maxNum+1 {
				d := s.distance
				if retZero {
					d += m.distances[0][s.current]
				}
				if d < minDistance {
					minDistance = d
				}
				continue
			}
			for i := 0; i <= m.maxNum; i++ {
				if !s.visited[i] {
					newState := state{make(map[int]bool), s.distance + m.distances[i][s.current], i}
					for k, v := range s.visited {
						newState.visited[k] = v
					}
					newState.visited[i] = true
					newStates = append(newStates, newState)
				}
			}
		}
		states = newStates
	}

	return minDistance
}
