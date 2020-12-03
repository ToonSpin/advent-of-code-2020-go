package main

import (
	"bufio"
	"fmt"
	"os"
)

func getTreesForSlope(input [][]bool, slopeX int, slopeY int) int {
	height := len(input)
	width := len(input[0])
	x := 0
	numTrees := 0
	for y := 0; y < height; {
		if input[y][x] {
			numTrees++
		}
		x += slopeX
		x %= width
		y += slopeY
	}
	return numTrees
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := make([][]bool, 0)

	for scanner.Scan() {
		line := make([]bool, 0)
		for _, c := range scanner.Text() {
			var b bool
			if c == '#' {
				b = true
			} else {
				b = false
			}
			line = append(line, b)
		}
		input = append(input, line)
	}

	partOne := getTreesForSlope(input, 3, 1)

	partTwo := partOne
	partTwo *= getTreesForSlope(input, 1, 1)
	partTwo *= getTreesForSlope(input, 5, 1)
	partTwo *= getTreesForSlope(input, 7, 1)
	partTwo *= getTreesForSlope(input, 1, 2)

	fmt.Printf("Number of trees with slope (3, 1): %d\n", partOne)
	fmt.Printf("Number of trees with all slopes: %d\n", partTwo)
}
