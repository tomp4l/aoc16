package day20

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	ranges, err := parse(input)
	if err != nil {
		return
	}
	min, total := valid(ranges)
	p1 = strconv.Itoa(min)
	p2 = strconv.Itoa(total)

	return
}

type ipRange struct {
	from int
	to   int
}

func parse(input string) ([]ipRange, error) {
	ret := make([]ipRange, 0)
	for _, l := range strings.Split(input, "\n") {
		split := strings.Split(l, "-")
		if len(split) != 2 {
			return nil, fmt.Errorf("bad ip range format: %s", l)
		}
		from, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		ret = append(ret, ipRange{from, to})
	}

	return ret, nil
}

func valid(ranges []ipRange) (int, int) {
	slices.SortFunc(ranges, func(a, b ipRange) int {
		return a.from - b.from
	})

	minTo := 0
	total := 0
	to := ranges[0].to
	for _, r := range ranges {
		if r.from <= to+1 {
			if to < r.to {
				to = r.to
			}
		} else {
			if minTo == 0 {
				minTo = to + 1
			}
			total += r.from - to - 1
			to = r.to
		}
	}

	return minTo, total
}
