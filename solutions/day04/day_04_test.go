package day04

import (
	"testing"
)

var input = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

func TestDay4Part1Example(t *testing.T) {
	expected := 18
	result, _ := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay4Part2Example(t *testing.T) {
	expected := 9
	_, result := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
