package day05

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Solver(input string) (int, int) {
	partOne, partTwo := 0, 0

	re := regexp.MustCompile(`(\d+)\|(\d+)|.*,*`)
	matches := re.FindAllStringSubmatch(input, -1)

	pageOrder := make(map[string][]string)

	for _, match := range matches {
		if strings.Contains(match[0], "|") {
			pageOrder[match[1]] = append(pageOrder[match[1]], match[2])
		}
		if strings.Contains(match[0], ",") {
			pageUpdates := strings.Split(match[0], ",")
			if isValidOrder(pageOrder, pageUpdates) {
				partOne += getMiddle(pageUpdates)
			} else {
				orderUpdated := make([]string, 0, len(pageUpdates))
				orderUpdated = append(orderUpdated, pageUpdates...)
				slices.SortFunc(orderUpdated, func(a string, b string) int {
					before := pageOrder[a]
					if len(before) == 0 {
						return 0
					}

					if slices.Contains(before, b) {
						return -1
					}

					return 1
				})
				partTwo += getMiddle(orderUpdated)
			}
		}
	}

	return partOne, partTwo
}

func isValidOrder(before map[string][]string, updates []string) bool {
	for idx, update := range updates {
		// Check if any of the rules are actually after it
		for _, rule := range before[update] {
			if slices.Contains(updates[:idx], rule) {
				// This would be false because the number expected to be before was actually after
				return false
			}
		}
	}

	return true
}

func getMiddle(slice []string) int {
	middle, _ := strconv.Atoi(slice[len(slice)/2])
	return middle
}
