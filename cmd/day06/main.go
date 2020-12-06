package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Passenger struct {
	answers map[rune]int
}

func tallyAnswersForGroup(group []Passenger) (int, int) {
	answers := make(map[rune]int)
	for _, p := range group {
		for r, _ := range p.answers {
			count, _ := answers[r]
			answers[r] = count + 1
		}
	}
	countForPart2 := 0
	numPassengers := len(group)
	for _, count := range answers {
		if count == numPassengers {
			countForPart2++
		}
	}
	return len(answers), countForPart2
}

func newPassenger(s string) Passenger {
	theirAnswers := make(map[rune]int)
	for _, b := range s {
		count, _ := theirAnswers[b]
		theirAnswers[b] = count + 1
	}
	return Passenger{answers: theirAnswers}
}

func getInput(r io.Reader) [][]Passenger {
	b, _ := ioutil.ReadAll(r)
	rawInput := string(b)
	rawGroups := strings.Split(rawInput, "\n\n")
	groups := make([][]Passenger, 0)

	for _, rawGroup := range rawGroups {
		passengers := make([]Passenger, 0)
		for _, p := range strings.Split(rawGroup, "\n") {
			if len(p) > 0 {
				passengers = append(passengers, newPassenger(p))
			}
		}
		groups = append(groups, passengers)
	}

	return groups
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)

	part1Count := 0
	part2Count := 0
	for _, group := range input {
		p1, p2 := tallyAnswersForGroup(group)
		part1Count += p1
		part2Count += p2
	}
	fmt.Println("Count where ANY of the passengers answered:", part1Count)
	fmt.Println("Count where ALL of the passengers answered:", part2Count)
}
