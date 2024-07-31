package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	ins, err := parse(input)
	if err != nil {
		return
	}
	c := newComputer()
	c.runAll(ins)

	p1 = strconv.Itoa(c.regs[reg("a")])

	c = newComputer()
	c.regs[reg("c")] = 1
	c.runAll(ins)

	p2 = strconv.Itoa(c.regs[reg("a")])

	return
}

const (
	copyPrefix = "cpy "
	incPrefix  = "inc "
	decPrefix  = "dec "
	jumpPrefix = "jnz "
)

func parse(input string) ([]instruction, error) {
	instructions := make([]instruction, 0)

	for _, l := range strings.Split(input, "\n") {
		var instruction instruction
		var err error
		switch {
		case strings.HasPrefix(l, copyPrefix):
			instruction, err = parseCopy(l)
		case strings.HasPrefix(l, incPrefix):
			instruction, err = parseIncrease(l)
		case strings.HasPrefix(l, decPrefix):
			instruction, err = parseDecrease(l)
		case strings.HasPrefix(l, jumpPrefix):
			instruction, err = parseJump(l)
		default:
			err = fmt.Errorf("unknown instruction: %s", l)
		}
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

type reg string
type computer struct {
	p    int
	regs map[reg]int
}

func newComputer() *computer {
	c := new(computer)

	c.regs = make(map[reg]int)
	return c
}

func (c *computer) runAll(ins []instruction) {
	for c.p < len(ins) {
		ins[c.p].execute(c)
	}
}

type instruction interface {
	execute(*computer)
}

type copy struct {
	i    int
	r    reg
	dest reg
}
type increase struct {
	reg reg
}
type decrease struct {
	reg reg
}
type jump struct {
	i      int
	reg    reg
	amount int
}

func (c *copy) execute(comp *computer) {
	if c.r != reg("") {
		comp.regs[c.dest] = comp.regs[c.r]
	} else {
		comp.regs[c.dest] = c.i
	}
	comp.p++
}
func (i *increase) execute(comp *computer) {
	comp.regs[i.reg]++
	comp.p++
}
func (d *decrease) execute(comp *computer) {
	comp.regs[d.reg]--
	comp.p++
}
func (j *jump) execute(comp *computer) {
	if j.i != 0 || comp.regs[j.reg] != 0 {
		comp.p += j.amount
	} else {
		comp.p++
	}
}

func regOrI(s string) (reg, int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return reg(s), 0
	} else {
		return reg(""), i
	}
}

func parseCopy(input string) (*copy, error) {
	rest := strings.TrimPrefix(input, copyPrefix)
	split := strings.Split(rest, " ")
	if len(split) != 2 {
		return nil, fmt.Errorf("bad copy fmt: %s", input)
	}
	c := new(copy)
	c.r, c.i = regOrI(split[0])
	c.dest = reg(split[1])
	return c, nil
}

func parseIncrease(input string) (*increase, error) {
	rest := strings.TrimPrefix(input, incPrefix)

	return &increase{reg(rest)}, nil
}

func parseDecrease(input string) (*decrease, error) {
	rest := strings.TrimPrefix(input, decPrefix)

	return &decrease{reg(rest)}, nil
}

func parseJump(input string) (*jump, error) {
	rest := strings.TrimPrefix(input, jumpPrefix)
	split := strings.Split(rest, " ")
	if len(split) != 2 {
		return nil, fmt.Errorf("bad jump fmt: %s", input)
	}
	j := new(jump)
	j.reg, j.i = regOrI(split[0])

	i, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, err
	}
	j.amount = i

	return j, nil
}
