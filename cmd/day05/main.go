package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getSeatId(boardingPass string) int {
	row := 0
	column := 0
	for i := 0; i < 7; i++ {
		row *= 2
		if boardingPass[i] == 'B' {
			row++
		}
	}

	for i := 7; i < 10; i++ {
		column *= 2
		if boardingPass[i] == 'R' {
			column++
		}
	}

	return row*8 + column
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxSeatId := 0
	seatIds := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		seatId := getSeatId(line)
		seatIds = append(seatIds, seatId)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}

	fmt.Println("The greatest seat ID among the boarding passes:", maxSeatId)

	prevId := -1
	sort.Ints(seatIds)
	for _, id := range seatIds {
		if id-prevId > 1 && prevId >= 0 {
			fmt.Println("My seat ID:", id-1)
			break
		}
		prevId = id
	}
}
