package day18

import "strconv"

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	p1 = strconv.Itoa(countSafe(extend(parse(input), 40)))
	p2 = strconv.Itoa(countSafe(extend(parse(input), 400000)))

	return
}

func parse(input string) (b []bool) {
	for _, c := range input {
		b = append(b, c == '.')
	}
	return
}

func extend(start []bool, length int) [][]bool {
	last := start
	ret := [][]bool{start}

	for len(ret) < length {
		next := make([]bool, len(last))

		next[0] = !isTrap(false, !last[0], !last[1])
		next[len(next)-1] = !isTrap(!last[len(next)-2], !last[len(next)-1], false)

		for i := 1; i < len(last)-1; i++ {
			l := !last[i-1]
			c := !last[i]
			r := !last[i+1]

			if !isTrap(l, c, r) {
				next[i] = true
			}
		}

		last = next
		ret = append(ret, next)
	}

	return ret
}

func isTrap(l, c, r bool) bool {
	return (l && c && !r) ||
		(c && r && !l) ||
		(l && !c && !r) ||
		(r && !c && !l)
}

func countSafe(cells [][]bool) (count int) {
	for _, r := range cells {
		for _, c := range r {
			if c {
				count++
			}
		}
	}
	return
}
