package day11

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type Day struct{}

func (Day) Run(input string) (p1 string, p2 string, err error) {
	s, err := parse(input)
	if err != nil {
		return
	}
	p1 = strconv.Itoa(solve(s))

	s2, err := parse(input)
	if err != nil {
		return
	}
	s2.floor1.objects["elerium"] = chipGenerator{true, true}
	s2.floor1.objects["dilithium"] = chipGenerator{true, true}
	p2 = strconv.Itoa(solve(s2))
	return
}

type floor struct {
	objects map[string]chipGenerator
}

type chipGenerator struct {
	generator bool
	chip      bool
}

func (cg *chipGenerator) exists() bool {
	return cg.generator || cg.chip
}

type state struct {
	elevator int
	floor1   *floor
	floor2   *floor
	floor3   *floor
	floor4   *floor
}

func (f *floor) configuration() string {
	gens := 0
	chips := 0
	for _, cg := range f.objects {
		if cg.generator {
			gens += 1
		}
		if cg.chip {
			chips += 1
		}
	}

	return fmt.Sprintf("G%dM%d", gens, chips)
}

func (s *state) configuration() string {
	return fmt.Sprintf("E%d%s%s%s%s", s.elevator,
		s.floor1.configuration(), s.floor2.configuration(), s.floor3.configuration(), s.floor4.configuration())
}

func parse(input string) (s *state, err error) {
	s = new(state)
	s.elevator = 1

	for _, l := range strings.Split(input, "\n") {
		words := strings.Split(l, " ")
		floor := &floor{make(map[string]chipGenerator)}
		switch words[1] {
		case "first":
			s.floor1 = floor
		case "second":
			s.floor2 = floor
		case "third":
			s.floor3 = floor
		case "fourth":
			s.floor4 = floor
		default:
			err = fmt.Errorf("unknown floor %s", words[1])
			return
		}

		for i := 0; i < len(words); i++ {
			var element string
			var isGenerator bool
			if strings.Contains(words[i], "generator") {
				isGenerator = true
				element = words[i-1]
			}
			if strings.Contains(words[i], "microchip") {
				element = strings.TrimSuffix(words[i-1], "-compatible")
			}
			if element != "" {
				cg := floor.objects[element]
				if isGenerator {
					cg.generator = true
				} else {
					cg.chip = true
				}
				floor.objects[element] = cg
			}
		}
	}

	return
}

func (s *state) isFinished() bool {
	return len(s.floor1.objects) == 0 &&
		len(s.floor2.objects) == 0 &&
		len(s.floor3.objects) == 0
}

func (s *state) isValid() bool {
	return s.floor1.isValid() &&
		s.floor2.isValid() &&
		s.floor3.isValid() &&
		s.floor4.isValid()
}

func (f *floor) isValid() bool {
	var hasGenerator, hasUnprotectedChip bool
	for _, v := range f.objects {
		if v.generator {
			hasGenerator = true
		} else {
			hasUnprotectedChip = hasUnprotectedChip || v.chip
		}
		if hasUnprotectedChip && hasGenerator {
			return false
		}
	}

	return true
}

func (s *state) floor(i int) *floor {
	switch i {
	case 1:
		return s.floor1
	case 2:
		return s.floor2
	case 3:
		return s.floor3
	case 4:
		return s.floor4
	default:
		return nil
	}
}

func (s *state) current() *floor {
	return s.floor(s.elevator)
}

func (s *state) up() *floor {
	return s.floor(s.elevator + 1)
}

func (s *state) down() *floor {
	return s.floor(s.elevator - 1)
}

func unsetCurrent(s *state, k string, chip bool) {
	c := s.current().objects[k]
	if chip {
		c.chip = false
	} else {
		c.generator = false
	}
	if c.exists() {
		s.current().objects[k] = c
	} else {
		delete(s.current().objects, k)
	}
}

func setFloor(f *floor, k string, chip bool) {
	c := f.objects[k]
	if chip {
		c.chip = true
	} else {
		c.generator = true
	}
	f.objects[k] = c
}

func (s *state) moveUpOne(k string, c bool) *state {
	return s.moveUpTwo(k, k, c, c)
}

func (s *state) moveDownOne(k string, c bool) *state {
	return s.moveDownTwo(k, k, c, c)
}

func (s *state) moveDownTwo(k1, k2 string, c1, c2 bool) *state {
	sn := *s
	switch s.elevator {
	case 1:
		return nil
	case 2:
		sn.floor1 = &floor{maps.Clone(s.floor1.objects)}
		sn.floor2 = &floor{maps.Clone(s.floor2.objects)}
	case 3:
		sn.floor2 = &floor{maps.Clone(s.floor2.objects)}
		sn.floor3 = &floor{maps.Clone(s.floor3.objects)}
	case 4:
		sn.floor3 = &floor{maps.Clone(s.floor3.objects)}
		sn.floor4 = &floor{maps.Clone(s.floor4.objects)}
	}
	s = &sn
	unsetCurrent(s, k1, c1)
	unsetCurrent(s, k2, c2)
	setFloor(s.down(), k1, c1)
	setFloor(s.down(), k2, c2)
	s.elevator = s.elevator - 1
	if s.isValid() {
		return s
	}
	return nil
}

func (s *state) moveUpTwo(k1, k2 string, c1, c2 bool) *state {
	sn := *s
	switch s.elevator {
	case 1:
		sn.floor1 = &floor{maps.Clone(s.floor1.objects)}
		sn.floor2 = &floor{maps.Clone(s.floor2.objects)}
	case 2:
		sn.floor2 = &floor{maps.Clone(s.floor2.objects)}
		sn.floor3 = &floor{maps.Clone(s.floor3.objects)}
	case 3:
		sn.floor3 = &floor{maps.Clone(s.floor3.objects)}
		sn.floor4 = &floor{maps.Clone(s.floor4.objects)}
	case 4:
		return nil
	}
	s = &sn
	unsetCurrent(s, k1, c1)
	unsetCurrent(s, k2, c2)
	setFloor(s.up(), k1, c1)
	setFloor(s.up(), k2, c2)

	s.elevator = s.elevator + 1
	if s.isValid() {
		return s
	}
	return nil
}

func solve(initial *state) int {
	seen := make(map[string]bool)
	openStates := []*state{initial}
	seen[initial.configuration()] = true

	i := 0
	for {
		if len(openStates) == 0 {
			panic("no states")
		}
		newOpenStates := make([]*state, 0)
		for _, s := range openStates {
			if s.isFinished() {
				return i
			}
			floor := s.current()

			movedBoth := false
			for k, v := range floor.objects {
				if v.chip {
					su := s.moveUpOne(k, true)
					newOpenStates = append(newOpenStates, su)
					sd := s.moveDownOne(k, true)
					newOpenStates = append(newOpenStates, sd)
				}
				if v.generator {
					su := s.moveUpOne(k, false)
					newOpenStates = append(newOpenStates, su)
					sd := s.moveDownOne(k, false)
					newOpenStates = append(newOpenStates, sd)
				}
				if !movedBoth && v.chip && v.generator {
					movedBoth = true
					su := s.moveUpTwo(k, k, true, false)
					newOpenStates = append(newOpenStates, su)
				}
				for k2, v2 := range floor.objects {
					if k != k2 {
						if v.chip && v2.chip {
							su := s.moveUpTwo(k, k2, true, true)
							newOpenStates = append(newOpenStates, su)
							sd := s.moveDownTwo(k, k2, true, true)
							newOpenStates = append(newOpenStates, sd)
						}
						if v.generator && v2.generator {
							su := s.moveUpTwo(k, k2, false, false)
							newOpenStates = append(newOpenStates, su)
							sd := s.moveDownTwo(k, k2, false, false)
							newOpenStates = append(newOpenStates, sd)
						}
					}
				}
			}
		}

		openStates = make([]*state, 0)
		for _, s := range newOpenStates {
			if s == nil {
				continue
			}
			if seen[s.configuration()] {
				continue
			}
			openStates = append(openStates, s)
			seen[s.configuration()] = true
		}
		i++
	}
}
