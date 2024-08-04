package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tomp4l/aoc16/day1"
	"github.com/tomp4l/aoc16/day2"
	"github.com/tomp4l/aoc16/day3"
	"github.com/tomp4l/aoc16/day4"
	"github.com/tomp4l/aoc16/day5"
	"github.com/tomp4l/aoc16/day6"
	"github.com/tomp4l/aoc16/day7"
	"github.com/tomp4l/aoc16/day8"
	"github.com/tomp4l/aoc16/day9"

	"github.com/tomp4l/aoc16/day10"
	"github.com/tomp4l/aoc16/day11"
	"github.com/tomp4l/aoc16/day12"
	"github.com/tomp4l/aoc16/day13"
	"github.com/tomp4l/aoc16/day14"
	"github.com/tomp4l/aoc16/day15"
	"github.com/tomp4l/aoc16/day16"
	"github.com/tomp4l/aoc16/day17"
	"github.com/tomp4l/aoc16/day18"
	"github.com/tomp4l/aoc16/day19"
	"github.com/tomp4l/aoc16/day20"
	"github.com/tomp4l/aoc16/day21"
)

func main() {
	days := map[int]Day{
		1:  day1.Day{},
		2:  day2.Day{},
		3:  day3.Day{},
		4:  day4.Day{},
		5:  day5.Day{},
		6:  day6.Day{},
		7:  day7.Day{},
		8:  day8.Day{},
		9:  day9.Day{},
		10: day10.Day{},
		11: day11.Day{},
		12: day12.Day{},
		13: day13.Day{},
		14: day14.Day{},
		15: day15.Day{},
		16: day16.Day{},
		17: day17.Day{},
		18: day18.Day{},
		19: day19.Day{},
		20: day20.Day{},
		21: day21.Day{},
	}

	args := os.Args

	if len(args) < 2 {
		log.Fatalf("Must provide a day")
	}

	day, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatalf("Failed to parse day: %v", err)
	}

	program, ok := days[day]
	if !ok {
		log.Fatalf("Undefined day: %d", day)
	}

	filename := "inputs/day%d.txt"
	input, err := os.ReadFile(fmt.Sprintf(filename, day))

	if err != nil {
		log.Fatalf("Failed to read input at %s: %v", filename, err)
	}

	part1, part2, err := program.Run(string(input))
	if err != nil {
		log.Fatalf("Day %d failed: %v", day, err)
	}

	fmt.Printf("Part 1: %s\n", part1)
	fmt.Printf("Part 2: %s\n", part2)
}

type Day interface {
	Run(string) (string, string, error)
}
