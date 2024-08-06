package computer

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	copyPrefix = "cpy "
	incPrefix  = "inc "
	decPrefix  = "dec "
	jumpPrefix = "jnz "
	tglPrefix  = "tgl "
	outPrefix  = "out "
)

func ParseInstructions(input string) ([]Instruction, error) {
	instructions := make([]Instruction, 0)

	for _, l := range strings.Split(input, "\n") {
		var instruction Instruction
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
		case strings.HasPrefix(l, tglPrefix):
			instruction, err = parseToggle(l)
		case strings.HasPrefix(l, outPrefix):
			instruction, err = parseOut(l)
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
type iValue int
type computer struct {
	p    int
	regs map[reg]int
	ins  []Instruction
	out  chan int
	done bool
}

func NewComputer(ins []Instruction) *computer {
	c := new(computer)

	c.regs = make(map[reg]int)
	c.ins = make([]Instruction, len(ins))
	c.out = make(chan int)
	copy(c.ins, ins)
	return c
}

func (comp *computer) RunAll() {
	for comp.p < len(comp.ins) && !comp.done {
		optimised := false

		if comp.p < len(comp.ins)-5 {
			cpy, ok := comp.ins[comp.p].(*cpy)
			if !ok || cpy.dest.register() == nil {
				goto multCheckEnd
			}
			dec, a := comp.ins[comp.p+1].(*decrease)
			inc, b := comp.ins[comp.p+2].(*increase)
			if !a || !b {
				dec, a = comp.ins[comp.p+2].(*decrease)
				inc, b = comp.ins[comp.p+1].(*increase)
			}
			if !a || !b {
				goto multCheckEnd
			}
			if dec.reg != *cpy.dest.register() {
				goto multCheckEnd
			}

			jmp1, ok := comp.ins[comp.p+3].(*jump)
			if !ok || jmp1.amount.value(comp) != -2 ||
				jmp1.check.register() == nil || dec.reg != *jmp1.check.register() {
				goto multCheckEnd
			}
			dec2, ok := comp.ins[comp.p+4].(*decrease)
			if !ok {
				goto multCheckEnd
			}
			jmp2, ok := comp.ins[comp.p+5].(*jump)
			if !ok || jmp2.amount.value(comp) != -5 ||
				jmp2.check.register() == nil || dec2.reg != *jmp2.check.register() {
				goto multCheckEnd
			}

			comp.regs[inc.reg] += cpy.source.value(comp) * dec2.reg.value(comp)
			comp.regs[dec.reg] = 0
			comp.regs[dec2.reg] = 0
			optimised = true
			comp.p += 6
		}
	multCheckEnd:

		if !optimised {
			comp.ins[comp.p].execute(comp)
		}
	}
}

func (c *computer) Reg(r string) int {
	return c.regs[reg(r)]
}

func (c *computer) SetReg(r string, i int) {
	c.regs[reg(r)] = i
}

func (c *computer) Close() {
	c.done = true
}

func (c *computer) Out() chan int {
	return c.out
}

type Instruction interface {
	execute(*computer)
	toggle() Instruction
}

type regOrInt interface {
	value(*computer) int
	register() *reg
}

func (r reg) value(c *computer) int {
	return c.regs[r]
}

func (r reg) register() *reg {
	return &r
}

func (i iValue) value(*computer) int {
	return int(i)
}

func (iValue) register() *reg {
	return nil
}

type cpy struct {
	source regOrInt
	dest   regOrInt
}
type increase struct {
	reg reg
}
type decrease struct {
	reg reg
}
type jump struct {
	check  regOrInt
	amount regOrInt
}
type toggle struct {
	reg reg
}

type out struct {
	reg reg
}

func (c *cpy) execute(comp *computer) {
	dest := c.dest.register()
	comp.p++
	if dest == nil {
		return
	}
	comp.regs[*dest] = c.source.value(comp)
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
	v := j.check.value(comp)
	if v != 0 {
		comp.p += j.amount.value(comp)
	} else {
		comp.p++
	}
}

func (r *toggle) execute(comp *computer) {
	i := comp.p + comp.regs[r.reg]
	comp.p++
	if i >= len(comp.ins) {
		return
	}
	comp.ins[i] = comp.ins[i].toggle()
}

func (o *out) execute(comp *computer) {
	comp.out <- comp.regs[o.reg]
	comp.p++
}

func regOrI(s string) regOrInt {
	i, err := strconv.Atoi(s)
	if err != nil {
		return reg(s)
	} else {
		return iValue(i)
	}
}

func parseCopy(input string) (*cpy, error) {
	rest := strings.TrimPrefix(input, copyPrefix)
	split := strings.Split(rest, " ")
	if len(split) != 2 {
		return nil, fmt.Errorf("bad copy fmt: %s", input)
	}
	c := new(cpy)
	c.source = regOrI(split[0])
	c.dest = regOrI(split[1])
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
	j.check = regOrI(split[0])
	j.amount = regOrI(split[1])

	return j, nil
}

func parseToggle(input string) (*toggle, error) {
	rest := strings.TrimPrefix(input, tglPrefix)

	return &toggle{reg(rest)}, nil
}

func parseOut(input string) (*out, error) {
	rest := strings.TrimPrefix(input, outPrefix)

	return &out{reg(rest)}, nil
}

func (i *cpy) toggle() Instruction {
	return &jump{check: i.source, amount: i.dest}
}
func (i *increase) toggle() Instruction {
	return &decrease{i.reg}
}
func (i *decrease) toggle() Instruction {
	return &increase{i.reg}
}
func (i *jump) toggle() Instruction {
	return &cpy{source: i.check, dest: i.amount}
}
func (i *toggle) toggle() Instruction {
	return &increase{i.reg}
}
func (i *out) toggle() Instruction {
	panic("undefined behaviour")
}
