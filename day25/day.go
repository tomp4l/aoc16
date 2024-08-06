package day25

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

	i := 0
	for {
		if checkClock(ins, i) {
			p1 = strconv.Itoa(i)
			break
		}
		i++
	}

	p2 = "finished!"

	return
}

func checkClock(ins []computer.Instruction, i int) bool {
	c := computer.NewComputer(ins)
	defer c.Close()
	c.SetReg("a", i)
	go func() {
		c.RunAll()
	}()

	x := <-c.Out()
	var y int

	for i := 0; i < 20; i++ {
		y = <-c.Out()

		if x == 1 && y == 0 {
			x = y
			continue
		}
		if x == 0 && y == 1 {
			x = y
			continue
		}
		return false
	}

	return true
}
