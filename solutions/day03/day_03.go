package day03

import (
	"regexp"
	"strconv"
)

func Solver(input string) (int, int) {
	partOne, partTwo := 0, 0
	readMul := true
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(don't\(\)|do\(\))`)

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if match[0] == "do()" {
			readMul = true
		} else if match[0] == "don't()" {
			readMul = false
		} else {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			partOne += num1 * num2
			if readMul {
				partTwo += num1 * num2
			}
		}
	}

	return partOne, partTwo
}
