package day23

import (
	"strconv"

	"github.com/tomp4l/aoc16/computer"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	ins, err := computer.ParseInstructions(input)
	if err != nil {
		return
	}
	c := computer.NewComputer(ins)
	c.SetReg("a", 7)
	c.RunAll()

	p1 = strconv.Itoa(c.Reg("a"))

	c = computer.NewComputer(ins)
	c.SetReg("a", 12)
	c.RunAll()

	p2 = strconv.Itoa(c.Reg("a"))

	return
}
