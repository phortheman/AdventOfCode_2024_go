package day02

import (
	"bufio"
	"strconv"
	"strings"
)

// Safely remove the element at the specified index. Create a new slice leaving the input unmodified
func removeIndex[T any](idx int, input []T) []T {
	if idx < 0 || idx >= len(input) {
		return input
	}

	// Create a new slice to return. Odd behavior with append
	output := make([]T, 0, len(input)-1)
	output = append(output, input[:idx]...)
	output = append(output, input[idx+1:]...)

	return output
}

// Start the recursive call to see if the report is safe. Starts with the first index and unknown direction
func isSafe(input []string) bool {
	return _isSafe(0, 0, input)
}

/*
	Do not use. Call 'isSafe' instead

idx: the current index we are checking

direction: 1 means we are ascending. -1 means we are desending. 0 means we don't know yet

input: the report
*/
func _isSafe(idx, direction int, input []string) bool {
	// If we're at the end then it is safe
	if idx == len(input)-1 {
		return true
	}

	currentElement, _ := strconv.Atoi(input[idx])
	nextElement, _ := strconv.Atoi(input[idx+1])

	// Get the difference between the current element and the next one
	diff := nextElement - currentElement

	switch diff {
	// No difference is unsafe
	case 0:
		return false

	// At least 1 and at most 3 difference
	case 1, 2, 3:
		if direction == 0 {
			direction = 1
		}

	case -1, -2, -3:
		if direction == 0 {
			direction = -1
		}

	// Everything else is unsafe
	default:
		return false
	}

	// If we are suppose to increase and it is decreasing, it is unsafe
	if direction == 1 && diff < 0 {
		return false
	}

	// If we are suppose to decrease and it is increasing, it is unsafe
	if direction == -1 && diff > 0 {
		return false
	}

	// Move the index over one and pass the expected direction
	return _isSafe(idx+1, direction, input)
}

func Solver(input string) (int, int) {
	safeReports := 0
	safeReportsWithTolerance := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		if isSafe(report) {
			safeReports += 1
		} else {
			for i := range report {
				modifiedReport := removeIndex(i, report)
				if isSafe(modifiedReport) {
					safeReportsWithTolerance += 1
					break
				}
			}
		}
	}

	return safeReports, safeReports + safeReportsWithTolerance
}
