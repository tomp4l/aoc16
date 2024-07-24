package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	parsed, err := parse(input)
	if err != nil {
		return
	}

	basic := &basicKeypad{1, 1}
	p1 = trackCode(basic, parsed)

	advanced := &advancedKeypad{0, 2}
	p2 = trackCode(advanced, parsed)

	return
}

func parse(input string) ([][]Direction, error) {
	directions := make([][]Direction, 0)
	for _, line := range strings.Split(input, "\n") {
		directionLine := make([]Direction, 0)
		for _, c := range line {
			switch c {
			case 'U':
				directionLine = append(directionLine, Up)
			case 'D':
				directionLine = append(directionLine, Down)
			case 'L':
				directionLine = append(directionLine, Left)
			case 'R':
				directionLine = append(directionLine, Right)
			default:
				return nil, fmt.Errorf("unknown direction %v", c)
			}
		}
		directions = append(directions, directionLine)

	}
	return directions, nil
}

func trackCode(keypad keypad, directions [][]Direction) string {
	var code string

	for _, l := range directions {
		for _, d := range l {
			keypad.move(d)
		}

		digit := keypad.digit()
		if digit < 10 {
			code += strconv.Itoa(digit)
		} else {
			switch digit {
			case 10:
				code += "A"
			case 11:
				code += "B"
			case 12:
				code += "C"
			case 13:
				code += "D"
			}
		}
	}

	return code
}

type Direction = int

const (
	Up Direction = iota
	Down
	Right
	Left
)

type keypad interface {
	digit() int
	move(d Direction)
}

type basicKeypad struct {
	x int
	y int
}

func (k *basicKeypad) digit() int {
	return 1 + k.x + k.y*3
}

func (k *basicKeypad) move(d Direction) {
	switch d {
	case Up:
		if k.y > 0 {
			k.y -= 1
		}
	case Down:
		if k.y < 2 {
			k.y += 1
		}
	case Right:
		if k.x < 2 {
			k.x += 1
		}
	case Left:
		if k.x > 0 {
			k.x -= 1
		}
	}
}

type advancedKeypad struct {
	x int
	y int
}

func (k *advancedKeypad) digit() int {
	switch k.y {
	case 0:
		return k.x - 1
	case 1:
		return k.x + 1
	case 2:
		return k.x + 5
	case 3:
		return k.x + 9
	case 4:
		return k.x + 11
	}

	panic("unreachable")
}

func (k *advancedKeypad) min(i int) int {
	switch i {
	case 0, 4:
		return 2
	case 1, 3:
		return 1
	case 2:
		return 0
	}

	panic("unreachable")
}

func (k *advancedKeypad) max(i int) int {
	switch i {
	case 0, 4:
		return 2
	case 1, 3:
		return 3
	case 2:
		return 4
	}

	panic("unreachable")
}

func (k *advancedKeypad) move(d Direction) {
	switch d {
	case Up:
		if k.y > k.min(k.x) {
			k.y -= 1
		}
	case Down:
		if k.y < k.max(k.x) {
			k.y += 1
		}
	case Right:
		if k.x < k.max(k.y) {
			k.x += 1
		}
	case Left:
		if k.x > k.min(k.y) {
			k.x -= 1
		}
	}
}
