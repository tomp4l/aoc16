package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

const (
	magicLow  = 17
	magicHigh = 61
)

func (Day) Run(input string) (p1 string, p2 string, err error) {

	bots, err := parse(input)
	if err != nil {
		return
	}

	p1 = strconv.Itoa(bots.findComparer(magicLow, magicHigh))
	bots.runEnd()
	p2 = strconv.Itoa(bots.outputs[0] * bots.outputs[1] * bots.outputs[2])

	return
}

const (
	botPrefix   = "bot "
	valuePrefix = "value "
)

func parse(input string) (*bots, error) {
	bots := &bots{make(map[int]*bot), make(map[int]int)}

	for _, l := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(l, botPrefix):
			err := parseInstruction(bots, l)
			if err != nil {
				return nil, err
			}
		case strings.HasPrefix(l, valuePrefix):
			err := parseValue(bots, l)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unrecognised bot format: %s", l)
		}

	}

	return bots, nil
}

func parseValue(bots *bots, input string) error {
	parts := strings.Split(input, " ")
	if len(parts) != 6 {
		return fmt.Errorf("expected 6 words: %s", input)
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(parts[5])
	if err != nil {
		return err
	}

	bot := bots.getBot(id)
	bot.addChip(value)

	return nil
}

func parseInstruction(bots *bots, input string) error {
	parts := strings.Split(input, " ")
	if len(parts) != 12 {
		return fmt.Errorf("expected 12 words: %s", input)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	lowDest := parts[5]
	low, err := strconv.Atoi(parts[6])
	if err != nil {
		return err
	}

	highDest := parts[10]
	high, err := strconv.Atoi(parts[11])
	if err != nil {
		return err
	}

	bot := bots.getBot(id)

	bot.instruction = instruction{
		low: low, lowBot: lowDest == "bot",
		high: high, highBot: highDest == "bot",
	}

	return nil
}

type bots struct {
	bots    map[int]*bot
	outputs map[int]int
}

func (b *bots) getBot(id int) *bot {
	bot_ := b.bots[id]
	if bot_ == nil {
		bot_ = new(bot)
		b.bots[id] = bot_
	}
	return bot_
}

func (b *bots) step1() bool {
	for _, bot := range b.bots {
		if bot.isReady() {
			bot.execute(b)
			return true
		}
	}

	return false
}

func (b *bots) findComparer(low, high int) int {
	for {
		for id, b := range b.bots {
			if b.isComparing(low, high) {
				return id
			}
		}
		if !b.step1() {
			return -1
		}
	}
}

func (b *bots) runEnd() {
	var ran = true
	for ran {
		ran = false
		for _, bot := range b.bots {
			if bot.isReady() {
				bot.execute(b)
				ran = true
			}
		}
	}
}

type bot struct {
	pending     int
	low         int
	high        int
	instruction instruction
}

func (b *bot) addChip(value int) {
	if b.pending != 0 && b.high != 0 {
		panic(fmt.Sprintf("invalid bot state %+v", b))
	}

	if b.pending == 0 {
		b.pending = value
	} else if b.pending < value {
		b.low = b.pending
		b.high = value
		b.pending = 0
	} else {
		b.high = b.pending
		b.low = value
		b.pending = 0
	}
}

func (b *bot) isReady() bool {
	return b.high != 0 && b.low != 0
}

func (b *bot) execute(bots *bots) {
	if !b.isReady() {
		panic("bot must be ready")
	}

	if b.instruction.highBot {
		bots.getBot(b.instruction.high).addChip(b.high)
	} else {
		bots.outputs[b.instruction.high] = b.high
	}

	if b.instruction.lowBot {
		bots.getBot(b.instruction.low).addChip(b.low)
	} else {
		bots.outputs[b.instruction.low] = b.low
	}

	b.low = 0
	b.high = 0
}

func (b *bot) isComparing(low, high int) bool {
	return b.low == low && b.high == high
}

type instruction struct {
	low     int
	lowBot  bool
	high    int
	highBot bool
}
