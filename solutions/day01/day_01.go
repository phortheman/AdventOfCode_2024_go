package day01

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func convStringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func Solver(input string) (int, int) {

	leftLocations := make([]int, 0, len(input))
	rightLocations := make([]int, 0, len(input))

	frequency := make(map[int]int)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		locationIDs := strings.Split(line, "   ")

		leftLocations = append(leftLocations, convStringToInt(locationIDs[0]))
		rightLocations = append(rightLocations, convStringToInt(locationIDs[1]))

		frequency[rightLocations[len(rightLocations)-1]] += 1
	}

	slices.Sort(leftLocations)
	slices.Sort(rightLocations)

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
