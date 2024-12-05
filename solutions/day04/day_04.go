package day04

import (
	"bytes"
)

func Solver(input []byte) (int, int) {
	partOne, partTwo := 0, 0
	grid := bytes.Split(input, []byte("\n"))

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// If 'X' try to find the pattern in all directions
			if grid[row][col] == 'X' {
				for dir := range DIRECTION {
					if searchXMAS(grid, dir, row, col) {
						partOne += 1
					}
				}
			}

			// If 'A' we assume it is the center of "MAS" and try to find it
			if grid[row][col] == 'A' && searchMAS(grid, row, col) {
				partTwo += 1
			}
		}
	}

	return partOne, partTwo
}

var DIRECTION = map[string][]int{
	"NORTH":      {-1, 0},
	"NORTH EAST": {-1, +1},
	"NORTH WEST": {-1, -1},
	"SOUTH":      {+1, 0},
	"SOUTH EAST": {+1, +1},
	"SOUTH WEST": {+1, -1},
	"EAST":       {0, +1},
	"WEST":       {0, -1},
}

const XMAS = "XMAS"

// Get the character at a specific spot in the grid safely. Return a null byte if it is out of bounds
func getChar(grid [][]byte, row, col int) byte {
	// Check for out of bounds
	if row < 0 || row >= len(grid) {
		return byte(0)
	}

	if col < 0 || col >= len(grid[row]) {
		return byte(0)
	}

	return grid[row][col]
}

// Search for the word "XMAS" in the specified direction
func searchXMAS(grid [][]byte, direction string, row, col int) bool {
	for idx := 0; idx < len(XMAS); idx++ {
		curRow := row + (idx * DIRECTION[direction][0])
		curCol := col + (idx * DIRECTION[direction][1])
		if getChar(grid, curRow, curCol) != XMAS[idx] {
			return false
		}
	}
	return true
}

// Search in a cross pattern for MAS. The current row/col is expected to already be 'A'
func searchMAS(grid [][]byte, row, col int) bool {
	if !((getChar(grid, row-1, col-1) == 'M' && getChar(grid, row+1, col+1) == 'S') ||
		(getChar(grid, row-1, col-1) == 'S' && getChar(grid, row+1, col+1) == 'M')) {
		return false
	}

	if !((getChar(grid, row-1, col+1) == 'M' && getChar(grid, row+1, col-1) == 'S') ||
		(getChar(grid, row-1, col+1) == 'S' && getChar(grid, row+1, col-1) == 'M')) {
		return false
	}

	return true
}
