package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	rooms, err := parse(input)

	if err != nil {
		return
	}

	p1 = strconv.Itoa(validRooms(rooms))
	p2 = strconv.Itoa(findRoom(rooms, "northpole object storage"))

	return
}

type room struct {
	name     string
	checksum string
	number   int
}

func parseRoom(input string) (r room, err error) {
	split := strings.Split(input, "-")
	if len(split) < 2 {
		err = fmt.Errorf("not enough parts: %s", input)
		return
	}

	r.name = strings.Join(split[:len(split)-1], "-")
	last := split[len(split)-1]
	breakpoint := len(last) - 7
	r.number, err = strconv.Atoi(last[:breakpoint])
	if err != nil {
		return
	}
	r.checksum = last[breakpoint+1 : len(last)-1]
	return
}

func parse(input string) ([]room, error) {
	rooms := make([]room, 0)
	for _, l := range strings.Split(input, "\n") {
		r, err := parseRoom(l)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}

	return rooms, nil
}

func validRooms(rooms []room) int {
	numbers := 0
	for _, r := range rooms {
		if r.checksum == checksum(r.name) {
			numbers += r.number
		}
	}
	return numbers
}

func checksum(name string) string {
	type pair struct {
		r rune
		i int
	}

	frequencies := make(map[rune]int)
	for _, c := range name {
		if c != '-' {
			frequencies[c] += 1
		}
	}

	sorted := make([]pair, 0, len(frequencies))
	for k, v := range frequencies {
		sorted = append(sorted, pair{k, v})
	}
	slices.SortFunc(sorted, func(a, b pair) int {
		if a.i == b.i {
			return int(a.r - b.r)
		}
		return b.i - a.i
	})

	sum := ""
	for _, c := range sorted[:5] {
		sum += string(c.r)
	}

	return sum
}

func decrypt(r room) string {
	decrypted := ""
	shift := r.number % 26

	for _, c := range r.name {
		if c == '-' {
			decrypted += " "
		} else {
			n := c + rune(shift)
			if n > 'z' {
				n -= rune(26)
			}
			decrypted += string(n)
		}
	}

	return decrypted
}

func findRoom(rooms []room, t string) int {
	for _, r := range rooms {
		if decrypt(r) == t {
			return r.number
		}
	}
	return -1
}
