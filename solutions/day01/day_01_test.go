package day01

import (
	"testing"
)

const input string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1Example(t *testing.T) {
	expected := 11
	result, _ := Solver(input)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestDay1Part2Example(t *testing.T) {
	expected := 31
	_, result := Solver(input)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}
