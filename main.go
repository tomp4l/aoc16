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
)

func main() {
	days := map[int]Day{
		1: day1.Day{},
		2: day2.Day{},
		3: day3.Day{},
		4: day4.Day{},
		5: day5.Day{},
		6: day6.Day{},
		7: day7.Day{},
		8: day8.Day{},
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
