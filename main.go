package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/phortheman/AdventOfCode_2024_go/solutions/day01"
)

var (
	inputs = []string{
		"inputs/day_01.txt",
	}
)

var (
	specificDays = false
	days         = make([]bool, len(inputs))
)

func init() {
	// Dynamically add the flags as more solutions are provided
	for i := 0; i < len(days); i++ {
		flag.BoolVar(&days[i], fmt.Sprint(i+1), false, fmt.Sprintf("Run day %02d", i+1))
	}

	flag.Parse()

	// If any flags are provided then run only specific days
	specificDays = slices.Contains(days, true)
}

func main() {
	var day, part1, part2 int
	for i, input := range inputs {
		day++
		// If a specific day is specified and this is not one of those days, skip
		if specificDays && !days[i] {
			continue
		}

		content, err := os.ReadFile(input)
		if err != nil {
			fmt.Printf("\nMissing day %d input file. Make sure you are saving the puzzle input like this 'inputs/day_%02d.txt'\n", day, day)
			continue
		}

		switch day {
		case 1:
			part1, part2 = day01.Solver(string(content))
		}

		fmt.Printf("\nDay %d	Part 1: %d\n", day, part1)
		fmt.Printf("Day %d	Part 2: %d\n", day, part2)

		part1, part2 = 0, 0
	}
}