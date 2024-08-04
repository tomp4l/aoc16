package day21

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	ins, err := parse(input)
	if err != nil {
		return
	}

	p1 = runAll("abcdefgh", ins)
	p2 = reverseAll("fbgdceah", ins)

	return
}

type instruction interface {
	execute(string) string
	reverse(string) string
}

func parse(input string) ([]instruction, error) {
	ret := make([]instruction, 0)
	for _, l := range strings.Split(input, "\n") {
		var ins instruction
		var err error = fmt.Errorf("unmatched line: %s", l)
		switch {
		case strings.HasPrefix(l, "swap position"):
			ins, err = parseSwapPosition(l)
		case strings.HasPrefix(l, "swap letter"):
			ins, err = parseSwapLetter(l)
		case strings.HasPrefix(l, "rotate based on position of letter"):
			ins, err = parseRotatePositionOf(l)
		case strings.HasPrefix(l, "rotate"):
			ins, err = parseRotateDir(l)
		case strings.HasPrefix(l, "reverse"):
			ins, err = parseReverse(l)
		case strings.HasPrefix(l, "move"):
			ins, err = parseMove(l)
		}
		if err != nil {
			return nil, err
		}
		ret = append(ret, ins)
	}

	return ret, nil
}

func runAll(word string, instructions []instruction) string {
	for _, i := range instructions {
		word = i.execute(word)
	}

	return word
}

func reverseAll(word string, instructions []instruction) string {
	for i := len(instructions) - 1; i >= 0; i-- {
		word = instructions[i].reverse(word)
	}

	return word
}

type swapPosition struct {
	from int
	to   int
}
type swapLetter struct {
	from rune
	to   rune
}
type rotatePositionOf struct {
	letter rune
}
type rotateDir struct {
	left   bool
	amount int
}
type reverse struct {
	from int
	to   int
}
type move struct {
	from int
	to   int
}

func splitWords(line string, name string, expectedLength int) ([]string, error) {
	split := strings.Split(line, " ")
	if len(split) != expectedLength {
		return nil, fmt.Errorf("%s expects %d words, got %d: %s", name, expectedLength, len(split), line)
	}
	return split, nil
}

func parseSwapPosition(line string) (*swapPosition, error) {
	split, err := splitWords(line, "swap position", 6)
	if err != nil {
		return nil, err
	}
	from, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(split[5])
	if err != nil {
		return nil, err
	}

	return &swapPosition{from, to}, nil
}

func parseSwapLetter(line string) (*swapLetter, error) {
	split, err := splitWords(line, "swap letter", 6)
	if err != nil {
		return nil, err
	}
	from, _ := utf8.DecodeRune([]byte(split[2]))
	to, _ := utf8.DecodeRune([]byte(split[5]))

	return &swapLetter{from, to}, nil
}

func parseRotatePositionOf(line string) (*rotatePositionOf, error) {
	split, err := splitWords(line, "rotate position", 7)
	if err != nil {
		return nil, err
	}
	letter, _ := utf8.DecodeRune([]byte(split[6]))

	return &rotatePositionOf{letter}, nil
}

func parseRotateDir(line string) (*rotateDir, error) {
	split, err := splitWords(line, "rotate direction", 4)
	if err != nil {
		return nil, err
	}
	left := split[1] == "left"
	amount, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}

	return &rotateDir{left, amount}, nil
}

func parseReverse(line string) (*reverse, error) {
	split, err := splitWords(line, "reverse", 5)
	if err != nil {
		return nil, err
	}
	from, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(split[4])
	if err != nil {
		return nil, err
	}

	return &reverse{from, to}, nil
}

func parseMove(line string) (*move, error) {
	split, err := splitWords(line, "move", 6)
	if err != nil {
		return nil, err
	}
	from, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(split[5])
	if err != nil {
		return nil, err
	}

	return &move{from, to}, nil
}

func (i *swapPosition) execute(word string) string {

	bytes := []byte(word)
	bytes[i.from], bytes[i.to] = bytes[i.to], bytes[i.from]

	return string(bytes)
}

func (s *swapLetter) execute(word string) string {

	var from, to int
	for i, c := range word {
		if c == s.from {
			from = i
		}
		if c == s.to {
			to = i
		}
	}

	return (&swapPosition{from, to}).execute(word)
}

func (r *rotatePositionOf) execute(word string) string {
	var pos int
	for i, c := range word {
		if c == r.letter {
			pos = i
			break
		}
	}

	if pos >= 4 {
		pos++
	}

	return (&rotateDir{amount: 1 + pos}).execute(word)
}

func (r *rotateDir) execute(word string) string {
	copy := []byte(word)
	length := len(word)

	for i := 0; i < length; i++ {
		if r.left {
			copy[i] = word[(i+r.amount)%length]
		} else {
			copy[i] = word[(i-(r.amount%length)+length)%length]
		}
	}

	return string(copy)
}

func (r *reverse) execute(word string) string {
	copy := []byte(word)

	for i := 0; i <= r.to-r.from; i++ {
		copy[i+r.from] = word[r.to-i]
	}

	return string(copy)
}

func (i *move) execute(word string) string {
	if i.from < i.to {
		return word[:i.from] + word[i.from+1:i.to+1] + word[i.from:i.from+1] + word[i.to+1:]
	} else {
		return word[:i.to] + word[i.from:i.from+1] + word[i.to:i.from] + word[i.from+1:]
	}
}

func (i *swapPosition) reverse(word string) string {
	return i.execute(word)
}

func (i *swapLetter) reverse(word string) string {
	return i.execute(word)
}

func (i *rotatePositionOf) reverse(word string) string {
	l1 := &rotateDir{left: true, amount: 1}
	attempt := word
	for {
		if i.execute(attempt) == word {
			return attempt
		}
		attempt = l1.execute(attempt)
	}
}

func (i *rotateDir) reverse(word string) string {
	return (&rotateDir{!i.left, i.amount}).execute(word)
}

func (i *reverse) reverse(word string) string {
	return i.execute(word)
}

func (i *move) reverse(word string) string {
	return (&move{from: i.to, to: i.from}).execute(word)
}
