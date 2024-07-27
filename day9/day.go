package day9

import (
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	decompressed := decompress(input)
	p1 = strconv.Itoa(len(decompressed))
	p2 = strconv.Itoa(decompress2(input))

	return
}

func decompress(s string) string {
	parts := make([]string, 0)
	last := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			parts = append(parts, s[last:i])
			var amount, length int
			amount, length = parseAmountLength(s, &i)

			toCopy := s[i : i+length]
			for j := 0; j < amount; j++ {
				parts = append(parts, toCopy)
			}
			i += length
			last = i
		} else {
			i++
		}
	}

	parts = append(parts, s[last:])

	return strings.Join(parts, "")
}

func decompress2(s string) int {
	length := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			var amount, l int
			amount, l = parseAmountLength(s, &i)

			expanded := decompress2(s[i : i+l])

			length += amount * expanded

			i += l
		} else {
			i++
			length++
		}
	}
	return length
}

func parseAmountLength(s string, i *int) (amount int, length int) {
	num1 := ""
	*i++
	for s[*i] != 'x' {
		num1 += string(s[*i])
		*i++
	}
	*i++
	num2 := ""
	for s[*i] != ')' {
		num2 += string(s[*i])
		*i++
	}
	*i++
	length, err := strconv.Atoi(num1)
	if err != nil {
		panic(err)
	}
	amount, err = strconv.Atoi(num2)
	if err != nil {
		panic(err)
	}
	return
}
