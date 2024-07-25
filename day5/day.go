package day5

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {

	p1 = password(input)
	p2 = password2(input)

	return
}

func password(input string) string {
	ret := ""

	for i := 0; len(ret) < 8; i++ {
		hash, pre := hasPrefix(input, i)
		if pre {
			ret += hash[5:6]
		}
	}

	return ret
}

func password2(input string) string {
	m := make(map[int]string, 0)

	for i := 0; len(m) < 8; i++ {
		hash, pre := hasPrefix(input, i)
		if pre {
			d, e := strconv.Atoi(hash[5:6])
			if d < 8 && e == nil {
				if _, ok := m[d]; !ok {
					m[d] = hash[6:7]
				}

			}

		}
	}

	ret := ""
	for i := 0; i < 8; i++ {
		ret += m[i]
	}

	return ret
}

func hasPrefix(input string, i int) (string, bool) {
	h := md5.New()
	io.WriteString(h, input+strconv.Itoa(i))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash, strings.HasPrefix(hash, "00000")
}
