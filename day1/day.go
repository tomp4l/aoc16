package day1

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tomp4l/aoc16/helpers"
)

type Day struct{}

type Direction int

const (
	Right Direction = iota
	Left
)

func (Day) Run(input string) (part1 string, part2 string, err error) {
	split := helpers.CommaSeperatedDay(input)

	part1i, err := distance(split)
	if err != nil {
		return
	}

	part2i, err := visitedTwice(split)
	if err != nil {
		return
	}

	part1 = strconv.Itoa(part1i)
	part2 = strconv.Itoa(part2i)

	return
}

func distance(directions []string) (int, error) {
	v, err := visited(directions)
	if err != nil {
		return 0, err
	}

	last := v[len(v)-1]

	return last.distance(), nil
}

func visitedTwice(directions []string) (int, error) {
	allVisited, err := visited(directions)
	if err != nil {
		return 0, err
	}

	var alreadyVisited = make(map[coord]bool)

	for _, v := range allVisited {
		if alreadyVisited[v] {
			return v.distance(), nil
		}
		alreadyVisited[v] = true
	}

	return 0, errors.New("didn't find twice visited")
}

type coord struct {
	x int
	y int
}

func (c *coord) distance() int {
	totalX := c.x
	totalY := c.y

	if totalX < 0 {
		totalX = -totalX
	}

	if totalY < 0 {
		totalY = -totalY
	}

	return totalX + totalY
}

func visited(directions []string) ([]coord, error) {
	var totalX = 0
	var totalY = 0
	var facing = 0
	coords := make([]coord, 0)

	for _, d := range directions {
		dir, dis, err := parseDirection(d)

		if err != nil {
			return nil, err
		}

		switch dir {
		case Right:
			facing += 1
		case Left:
			facing -= 1
		}

		facing += 4
		facing %= 4

		for i := 0; i < dis; i++ {
			switch facing {
			case 0:
				totalY += 1
			case 1:
				totalX += 1
			case 2:
				totalY -= 1
			case 3:
				totalX -= 1
			}
			coords = append(coords, coord{totalX, totalY})
		}
	}

	return coords, nil
}

func parseDirection(s string) (d Direction, distance int, err error) {

	if len(s) < 2 {
		err = fmt.Errorf("too short direction string %s", s)
		return
	}
	switch string(s[0]) {
	case "R":
		d = Right
	case "L":
		d = Left
	default:
		err = fmt.Errorf("unknown direction %s", s)
		return
	}

	distance, err = strconv.Atoi(string(s[1:]))
	return
}
