package day01

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func sortInsert(data []int, val int) []int {
	idx := sort.SearchInts(data, val)
	if idx == len(data) {
		return append(data, val)
	}

	// Make room
	data = append(data[:idx+1], data[idx:]...)

	data[idx] = val
	return data
}

func Solver(input string) (int, int) {
	leftLocations := make([]int, 0, len(input))
	rightLocations := make([]int, 0, len(input))
	frequency := make(map[int]int)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		locationIDs := strings.Split(line, "   ")

		left, _ := strconv.Atoi(locationIDs[0])
		right, _ := strconv.Atoi(locationIDs[1])
		frequency[right] += 1

		leftLocations = sortInsert(leftLocations, left)
		rightLocations = sortInsert(rightLocations, right)
	}

	distance := 0
	similarity := 0
	for i := 0; i < len(leftLocations); i++ {
		diff := leftLocations[i] - rightLocations[i]

		if diff < 0 {
			diff *= -1
		}

		distance += diff
		similarity += leftLocations[i] * frequency[leftLocations[i]]
	}

	return distance, similarity
}
