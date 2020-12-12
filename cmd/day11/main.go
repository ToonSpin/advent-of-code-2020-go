package main

import (
	"bufio"
	"fmt"
	"os"
)

func constrain(i, min, max int) int {
	if i < min {
		return min
	}

	if i > max {
		return max
	}

	return i
}

func countOccupiedNeighbors(input [][]rune, x, y int, partTwo bool) int {
	numOccupied := 0

	width := len(input[0])
	height := len(input)

	for dy := -1; dy <= 1; dy++ {
	outer:
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			p := x
			q := y

			for {
				p += dx
				q += dy

				if p < 0 || q < 0 {
					continue outer
				}
				if p >= width || q >= height {
					continue outer
				}

				if input[q][p] == '#' {
					numOccupied++
					break
				}
				if input[q][p] == 'L' {
					continue outer
				}

				if !partTwo {
					break
				}
			}
		}
	}

	return numOccupied
}

func iterate(input [][]rune, partTwo bool) (int, bool) {
	numOccupied := 0
	changed := false
	var occupiedThreshold int
	if partTwo {
		occupiedThreshold = 5
	} else {
		occupiedThreshold = 4
	}

	width := len(input[0])
	height := len(input)

	newInput := make([][]rune, height)
	for y, row := range input {
		newInput[y] = make([]rune, width)
		copy(newInput[y], input[y])

		counts := make([]int, 0)
		for x, seat := range row {
			occupiedNeighbors := countOccupiedNeighbors(input, x, y, partTwo)
			counts = append(counts, occupiedNeighbors)

			newSeat := seat
			if seat == 'L' && occupiedNeighbors == 0 {
				newSeat = '#'
				changed = true
			}
			if seat == '#' && occupiedNeighbors >= occupiedThreshold {
				newSeat = 'L'
				changed = true
			}
			if newSeat == '#' {
				numOccupied++
			}

			newInput[y][x] = newSeat
		}
	}

	for y, _ := range newInput {
		copy(input[y], newInput[y])
	}
	return numOccupied, changed
}

func getOccupiedInEquilibrium(input [][]rune, partTwo bool) int {
	numOccupied := 0
	changed := true

	for {
		numOccupied, changed = iterate(input, partTwo)
		if !changed {
			break
		}
	}
	return numOccupied
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := make([][]rune, 0)
	input2 := make([][]rune, 0)

	for scanner.Scan() {
		line := make([]rune, 0)
		for _, c := range scanner.Text() {
			line = append(line, c)
		}
		input = append(input, line)
		input2 = append(input2, line)
	}

	fmt.Println("NUmber of occupied seats (part 1):", getOccupiedInEquilibrium(input, false))
	fmt.Println("NUmber of occupied seats (part 2):", getOccupiedInEquilibrium(input2, true))
}
