package day05

import (
	"testing"
)

const input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDay5Part1Example(t *testing.T) {
	expected := 143
	result, _ := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay5Part2Example(t *testing.T) {
	expected := 123
	_, result := Solver(input)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay5GetMiddle(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"75", "47", "61", "53", "29"}, 61},
		{[]string{"97", "61", "53", "29", "13"}, 53},
		{[]string{"75", "29", "13"}, 29},
		{[]string{"75", "97", "47", "61", "53"}, 47},
		{[]string{"61", "13", "29"}, 13},
		{[]string{"97", "13", "75", "29", "47"}, 75},
	}

	for idx, test := range tests {
		result := getMiddle(test.input)
		if result != test.expected {
			t.Errorf("Test #%d expected %d but got %d: %v", idx+1, test.expected, result, test.input)
		}
	}
}
