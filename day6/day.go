package day6

import "strings"

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	lines := strings.Split(input, "\n")

	p1 = corrected(lines, true)
	p2 = corrected(lines, false)

	return
}

func corrected(input []string, most bool) string {
	freqs := make(map[int]map[rune]int)

	for _, in := range input {
		for i, c := range in {
			f := freqs[i]
			if f == nil {
				f = make(map[rune]int)
				freqs[i] = f
			}
			f[c] += 1
		}
	}

	ret := ""
	for i := 0; i < len(freqs); i++ {
		f := freqs[i]
		var min_max int
		if !most {
			min_max = len(input) + 1
		}
		var c string
		for k, v := range f {
			if most && v > min_max || !most && v < min_max {
				min_max = v
				c = string(k)
			}
		}
		ret += c
	}

	return ret
}
