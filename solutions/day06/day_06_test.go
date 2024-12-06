package day06

import (
	"testing"
)

const input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6Part1Example(t *testing.T) {
	expected := 41
	result, _ := Solver([]byte(input))

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay6Part2Example(t *testing.T) {
	expected := 6
	_, result := Solver([]byte(input))

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
