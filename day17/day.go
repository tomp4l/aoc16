package day17

import (
	"strconv"

	"github.com/tomp4l/aoc16/helpers"
)

type Day struct{}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (Day) Run(input string) (p1 string, p2 string, err error) {

	p1, longest := dfs(input)
	p2 = strconv.Itoa(len(longest))

	return
}

func path(moves []direction) string {
	path := ""
	for _, d := range moves {
		switch d {
		case up:
			path += "U"
		case down:
			path += "D"
		case left:
			path += "L"
		case right:
			path += "R"
		}
	}
	return path
}

func openDoors(passcode string, moves []direction) []direction {

	hash := helpers.Md5String(passcode + path(moves))

	dirs := make([]direction, 0)
	for i := 0; i < 4; i++ {
		if hash[i] > 'a' {
			var dir direction
			switch i {
			case 0:
				dir = up
			case 1:
				dir = down
			case 2:
				dir = left
			case 3:
				dir = right
			}
			dirs = append(dirs, dir)
		}
	}
	return dirs
}

func dfs(passcode string) (string, string) {

	type state struct {
		x       int
		y       int
		visited []direction
	}

	openStates := []state{
		{
			x:       0,
			y:       0,
			visited: make([]direction, 0),
		},
	}

	first := ""
	last := ""

	for len(openStates) > 0 {
		nextStates := make([]state, 0)
		for _, o := range openStates {
			if o.x == 3 && o.y == 3 {
				last = path(o.visited)
				if first == "" {
					first = last
				}
				continue
			}

			doors := openDoors(passcode, o.visited)

			for _, d := range doors {
				next := o

				switch d {
				case up:
					if next.y > 0 {
						next.y--
					}
				case down:
					if next.y < 3 {
						next.y++
					}
				case left:
					if next.x > 0 {
						next.x--
					}
				case right:
					if next.x < 3 {
						next.x++
					}
				}

				if next.x != o.x || next.y != o.y {
					visited := make([]direction, len(next.visited)+1)
					copy(visited, next.visited)
					visited[len(next.visited)] = d
					next.visited = visited
					nextStates = append(nextStates, next)
				}
			}
		}
		openStates = nextStates
	}

	return first, last
}
