package main

import (
	// "errors"
	"fmt"
	"sort"
)

func partOne(input []int) int {
	diff1 := 0
	diff3 := 0

	prev := 0
	for _, joltage := range input {
		diff := joltage - prev
		if diff == 1 {
			diff1++
		}
		if diff == 3 {
			diff3++
		}
		prev = joltage
	}

	return diff1 * diff3
}

func partTwo(input []int, index int, memo map[int]int) int {
	cached, ok := memo[index]
	if ok {
		return cached
	}

	if index >= len(input)-1 {
		return 1
	}

	total := 0
	first := input[index]
	for i := index + 1; i < len(input) && input[i]-first <= 3; i++ {
		total += partTwo(input, i, memo)
	}

	memo[index] = total
	return total
}

func getInput() []int {
	var input []int
	for {
		var entry int
		if _, err := fmt.Scanf("%d\n", &entry); err == nil {
			input = append(input, entry)
		} else {
			break
		}
	}

	input = append(input, 0)
	sort.Ints(input)

	max := input[len(input)-1]
	input = append(input, max+3)
	sort.Ints(input)

	return input
}

func main() {
	input := getInput()

	fmt.Println("The sum of the \"one\" and the \"three\" differences:", partOne(input))

	memo := make(map[int]int)
	fmt.Println("The number of possible combinations:", partTwo(input, 0, memo))
}
