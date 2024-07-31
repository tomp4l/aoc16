package day14

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	p1 = strconv.Itoa(key64(input, md5si))
	p2 = strconv.Itoa(key64(input, stretched))

	return
}

func key64(salt string, hasher func(string, int) string) int {
	index := 0

	type hit struct {
		r rune
		i int
	}
	keys := make(map[int]bool)
	hits := make([]hit, 0)

	end := 0

	for {
		h := hasher(salt, index)
		var l rune
		var count int
		first := true
		for _, c := range h {
			if c == l {
				count++
			} else {
				count = 1
				l = c
			}
			if count == 3 && first {
				first = false
				hits = append(hits, hit{c, index})
			}
			if count == 5 {
				for i := 0; i < len(hits); i++ {
					if hits[i].r == c && hits[i].i != index {
						keys[hits[i].i] = true
					}
				}

				if end == 0 && len(keys) >= 64 {
					end = index + 1000
				}
			}
		}
		index++

		if index == end {

			ks := make([]int, 0)

			for k := range keys {
				ks = append(ks, k)
			}

			sort.Ints(ks)

			return ks[63]
		}

		for i := len(hits) - 1; i >= 0; i-- {
			if index-hits[i].i > 1000 {
				hits = hits[i+1:]
				break
			}
		}
	}
}

func md5s(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func md5si(s string, i int) string {
	return md5s(s + strconv.Itoa(i))
}

func stretched(s string, i int) string {
	h := md5si(s, i)

	for i := 0; i < 2016; i++ {
		h = md5s(h)
	}

	return h
}
