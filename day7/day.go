package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	ips, err := parse(input)
	if err != nil {
		return
	}

	p1 = strconv.Itoa(countTls(ips))
	p2 = strconv.Itoa(countSsl(ips))

	return
}

type ip struct {
	supers []string
	hypers []string
}

func parse(input string) ([]ip, error) {
	ips := make([]ip, 0)
	for _, l := range strings.Split(input, "\n") {
		ip, err := parseIp(l)
		if err != nil {
			return nil, err
		}
		ips = append(ips, ip)
	}
	return ips, nil
}

func parseIp(in string) (ip, error) {
	s := in
	ip := ip{
		make([]string, 0),
		make([]string, 0),
	}
	for {
		b1 := strings.IndexRune(s, '[')
		b2 := strings.IndexRune(s, ']')
		if b1 == -1 {
			ip.supers = append(ip.supers, s)
			return ip, nil
		}
		if b2 == -1 {
			return ip, fmt.Errorf("unmatched brackets: %s", in)
		}
		ip.supers = append(ip.supers, s[:b1])
		ip.hypers = append(ip.hypers, s[b1+1:b2])

		s = s[b2+1:]
	}
}

func abba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		a1 := s[i]
		b1 := s[i+1]
		b2 := s[i+2]
		a2 := s[i+3]
		if a1 == a2 && b1 == b2 && a1 != b1 {
			return true
		}
	}
	return false
}

func (ip *ip) supportsTls() bool {
	for _, h := range ip.hypers {
		if abba(h) {
			return false
		}
	}

	for _, s := range ip.supers {
		if abba(s) {
			return true
		}
	}

	return false
}

func (ip *ip) supportsSsl() bool {
	type ab struct {
		a byte
		b byte
	}
	abas := make([]ab, 0)
	for _, s := range ip.supers {
		for i := 0; i < len(s)-2; i++ {
			a1 := s[i]
			b := s[i+1]
			a2 := s[i+2]
			if a1 == a2 && a1 != b {
				abas = append(abas, ab{a1, b})
			}
		}
	}

	for _, h := range ip.hypers {
		for i := 0; i < len(h)-2; i++ {
			b1 := h[i]
			a := h[i+1]
			b2 := h[i+2]
			for _, ab := range abas {
				if b1 == ab.b && a == ab.a && b2 == ab.b {
					return true
				}
			}
		}
	}

	return false
}

func countTls(ips []ip) int {
	count := 0
	for _, ip := range ips {
		if ip.supportsTls() {
			count++
		}
	}

	return count
}

func countSsl(ips []ip) int {
	count := 0
	for _, ip := range ips {
		if ip.supportsSsl() {
			count++
		}
	}

	return count
}
