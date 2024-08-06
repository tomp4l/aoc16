package day22

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	nodes, err := parse(input)
	if err != nil {
		return
	}

	p1 = strconv.Itoa(countViablePairs(nodes))

	grid := makeGrid(nodes)
	p2 = strconv.Itoa(shortestPrize(grid))

	return
}

type node struct {
	x         int
	y         int
	size      int
	used      int
	available int
}

func parseSize(size string) (int, error) {
	if size == "" {
		return 0, errors.New("empty size")
	}
	return strconv.Atoi(size[:len(size)-1])
}

func parseNode(line string) (*node, error) {
	split := strings.Split(line, " ")
	nonEmpty := make([]string, 0, 5)

	for _, s := range split {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}

	if len(nonEmpty) != 5 {
		return nil, fmt.Errorf("invalid node format: %s", line)
	}

	filename := nonEmpty[0]
	parts := strings.Split(filename, "-")
	if len(parts) != 3 || !strings.HasPrefix(parts[1], "x") || !strings.HasPrefix(parts[2], "y") {
		return nil, fmt.Errorf("bad filename: %s", filename)
	}

	node := new(node)
	var err error
	node.x, err = strconv.Atoi(parts[1][1:])
	if err != nil {
		return nil, err
	}
	node.y, err = strconv.Atoi(parts[2][1:])
	if err != nil {
		return nil, err
	}
	node.size, err = parseSize(nonEmpty[1])
	if err != nil {
		return nil, err
	}
	node.used, err = parseSize(nonEmpty[2])
	if err != nil {
		return nil, err
	}
	node.available, err = parseSize(nonEmpty[3])
	if err != nil {
		return nil, err
	}

	return node, nil
}

func parse(input string) ([]*node, error) {
	split := strings.Split(input, "\n")
	if len(split) < 2 {
		return nil, fmt.Errorf("expecting more than 2 lines: %s", input)
	}
	nodes := make([]*node, 0, len(split)-2)

	for _, l := range split[2:] {
		n, err := parseNode(l)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func countViablePairs(nodes []*node) int {
	count := 0
	for _, a := range nodes {
		if a.used == 0 {
			continue
		}
		for _, b := range nodes {
			if a != b {
				if a.used <= b.available {
					count++
				}
			}
		}
	}

	return count
}

type coord struct {
	x int
	y int
}

type grid struct {
	empty coord
	prize coord
	nodes map[coord]bool
}

func (g *grid) id() int {
	return g.empty.x |
		g.empty.y<<8 |
		g.prize.x<<16 |
		g.prize.y<<24
}

func makeGrid(nodes []*node) *grid {
	var empty *node
	var maxX *node
	for _, n := range nodes {
		if n.used == 0 {
			empty = n
		}
		if n.y == 0 && (maxX == nil || n.x > maxX.x) {
			maxX = n
		}
	}
	g := new(grid)
	g.empty.x = empty.x
	g.empty.y = empty.y
	g.prize.x = maxX.x
	g.nodes = make(map[coord]bool)

	for _, n := range nodes {
		if n.used <= empty.available {
			g.nodes[coord{n.x, n.y}] = true
		}
	}

	return g
}

func (g *grid) neighbours() []*grid {
	n := make([]*grid, 0)

	for i := 0; i < 4; i++ {
		c := g.empty
		switch i {
		case 0:
			c.x++
		case 1:
			c.x--
		case 2:
			c.y++
		case 3:
			c.y--
		}
		if g.nodes[c] {
			ng := new(grid)
			if c == g.prize {
				ng.prize = g.empty
			} else {
				ng.prize = g.prize
			}
			ng.empty = c
			ng.nodes = g.nodes
			n = append(n, ng)
		}
	}

	return n
}

func shortestPrize(g *grid) int {

	seen := make(map[int]bool)
	seen[g.id()] = true
	states := []*grid{g}
	steps := 0

	for {
		newStates := make([]*grid, 0)
		for _, g := range states {
			if g.prize == (coord{0, 0}) {
				return steps
			}
			newStates = append(newStates, g.neighbours()...)
		}

		states = make([]*grid, 0)
		for _, n := range newStates {
			if !seen[n.id()] {
				seen[n.id()] = true
				states = append(states, n)
			}
		}
		steps++
	}
}
