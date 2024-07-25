package day8

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	instructions, err := parse(input)
	if err != nil {
		return
	}

	screen := applyAll(instructions)
	p1 = strconv.Itoa(screen.countPixels())
	p2 = fmt.Sprintf("\n%s", screen)

	return
}

func parse(input string) ([]instruction, error) {
	ins := make([]instruction, 0)
	for _, l := range strings.Split(input, "\n") {
		i, err := parseInstruction(l)
		if err != nil {
			return nil, err
		}
		ins = append(ins, i)
	}

	return ins, nil
}

func applyAll(ins []instruction) *screen {
	s := newScreen()
	for _, i := range ins {
		i.apply(s)
	}
	return s
}

type screen struct {
	pixels [][]bool
}

const (
	screenWidth  = 50
	screenHeight = 6
)

func newScreen() *screen {
	screen := new(screen)
	screen.pixels = make([][]bool, screenHeight)
	for i := range screen.pixels {
		screen.pixels[i] = make([]bool, screenWidth)
	}
	return screen
}

func (s *screen) rect(a, b int) {
	for y := 0; y < b; y++ {
		for x := 0; x < a; x++ {
			s.pixels[y][x] = true
		}
	}
}

func (s *screen) rotateRow(y, b int) {
	row := make([]bool, screenWidth)
	copy(row, s.pixels[y])
	for i := 0; i < screenWidth; i++ {
		s.pixels[y][(i+b)%screenWidth] = row[i]
	}
}

func (s *screen) rotateColumn(x, b int) {
	column := make([]bool, screenHeight)
	for i := 0; i < screenHeight; i++ {
		column[i] = s.pixels[i][x]
	}
	for i := 0; i < screenHeight; i++ {
		s.pixels[(i+b)%screenHeight][x] = column[i]
	}
}

func (s *screen) countPixels() int {
	count := 0
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if s.pixels[y][x] {
				count++
			}
		}
	}
	return count
}

func (s *screen) String() string {
	ret := ""
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if s.pixels[y][x] {
				ret += "#"
			} else {
				ret += " "
			}
		}
		ret += "\n"
	}
	return ret
}

type instruction interface {
	apply(*screen)
}

const (
	rectPrefix      = "rect "
	rotateRowPrefix = "rotate row y="
	rotateColPrefix = "rotate column x="
)

func parseInstruction(input string) (instruction, error) {
	switch {
	case strings.HasPrefix(input, rectPrefix):
		return parseRect(input)
	case strings.HasPrefix(input, rotateColPrefix):
		return parseRotateColumn(input)
	case strings.HasPrefix(input, rotateRowPrefix):
		return parseRotateRow(input)
	}

	return nil, fmt.Errorf("unparseable instruction: %s", input)
}

func extract2(split []string) (int, int, error) {
	if len(split) != 2 {
		return 0, 0, fmt.Errorf("expecing two elements, got  %d", len(split))
	}
	var err error
	a, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}
	return a, b, nil
}

func parseRect(input string) (*rect, error) {
	rect := new(rect)
	split := strings.Split(strings.TrimPrefix(input, rectPrefix), "x")
	var err error
	rect.a, rect.b, err = extract2(split)
	return rect, err
}

type rect struct {
	a int
	b int
}

func (r *rect) apply(s *screen) {
	s.rect(r.a, r.b)
}

func parseRotateRow(input string) (*rotateRow, error) {
	rotateRow := new(rotateRow)
	split := strings.Split(strings.TrimPrefix(input, rotateRowPrefix), " by ")
	var err error
	rotateRow.y, rotateRow.b, err = extract2(split)
	return rotateRow, err
}

type rotateRow struct {
	y int
	b int
}

func (r *rotateRow) apply(s *screen) {
	s.rotateRow(r.y, r.b)
}

func parseRotateColumn(input string) (*rotateColumn, error) {
	rotateColumn := new(rotateColumn)
	split := strings.Split(strings.TrimPrefix(input, rotateColPrefix), " by ")
	var err error
	rotateColumn.x, rotateColumn.b, err = extract2(split)
	return rotateColumn, err
}

type rotateColumn struct {
	x int
	b int
}

func (r *rotateColumn) apply(s *screen) {
	s.rotateColumn(r.x, r.b)
}
