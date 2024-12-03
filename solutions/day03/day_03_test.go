package day03

import (
	"testing"
)

func TestDay3Part1Example(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	result, _ := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay3Part2Example(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48
	_, result := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
