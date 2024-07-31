package day16

import "fmt"

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	hd := newHd(input, 272)
	p1 = hd.checksum()

	hd = newHd(input, 35651584)
	p2 = hd.checksum()
	return
}

type hardDrive struct {
	size   int
	filled int
	bits   []bool
}

func newHd(init string, size int) *hardDrive {
	h := &hardDrive{
		bits:   make([]bool, size),
		size:   size,
		filled: len(init),
	}
	for i, c := range init {
		h.bits[i] = c == '1'
	}

	for h.size != h.filled {
		if h.size < h.filled {
			panic(fmt.Errorf("bad size hard drive (size smaller than init?) %d %d", h.size, h.filled))
		}
		h.fill()
	}
	return h
}

func (hd *hardDrive) fill() {
	p := 1
	for hd.filled+p < hd.size && p <= hd.filled {
		hd.bits[hd.filled+p] = !hd.bits[hd.filled-p]
		p++
	}
	hd.filled = hd.filled + p
}

func (hd *hardDrive) checksum() string {
	for hd.size%2 == 0 {
		nhd := new(hardDrive)
		nhd.size = hd.size / 2
		nhd.filled = nhd.size
		nhd.bits = make([]bool, nhd.size)
		for i := 0; i < nhd.size; i++ {
			nhd.bits[i] = hd.bits[2*i] == hd.bits[2*i+1]
		}
		hd = nhd
	}
	return hd.String()

}

func (hd *hardDrive) String() string {
	ret := ""
	for _, b := range hd.bits {
		if b {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}
