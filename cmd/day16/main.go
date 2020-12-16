package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ticket []int64

type Range struct {
	name string
	min1 int64
	max1 int64
	min2 int64
	max2 int64
}

func (r Range) validate(n int64) bool {
	if n >= r.min1 && n <= r.max1 {
		return true
	}
	if n >= r.min2 && n <= r.max2 {
		return true
	}
	return false
}

type FieldPositions struct {
	ranges        []Range
	possibilities []map[int]bool
}

func makeFieldPositions(ranges []Range) FieldPositions {
	possibilities := make([]map[int]bool, 0)
	numRanges := len(ranges)
	for i := 0; i < numRanges; i++ {
		p := make(map[int]bool)
		for j := 0; j < numRanges; j++ {
			p[j] = true
		}
		possibilities = append(possibilities, p)
	}
	return FieldPositions{ranges, possibilities}
}

func (p FieldPositions) updatePossibilities(ticket Ticket) {
	for rangeIndex, r := range p.ranges {
		for fieldIndex, v := range ticket {
			if !r.validate(v) {
				delete(p.possibilities[rangeIndex], fieldIndex)
				if len(p.possibilities[rangeIndex]) == 1 {
					p.clean()
				}
			}
		}
	}
}

func (p FieldPositions) purge(value int) bool {
	found := false
	for rangeIndex, poss := range p.possibilities {
		if len(poss) == 1 {
			continue
		}
		_, ok := poss[value]
		if ok {
			found = true
			delete(p.possibilities[rangeIndex], value)
		}
	}
	return found
}

func (p FieldPositions) clean() {
	found := true
	for found {
		found = false
		for _, poss := range p.possibilities {
			if len(poss) > 1 {
				continue
			}

			for value, _ := range poss {
				if p.purge(value) {
					found = true
				}
			}
		}
	}
}

func getRangesFromInput(input string) []Range {
	ranges := make([]Range, 0)
	rangeRe := regexp.MustCompile(`([^\n:]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, matches := range rangeRe.FindAllStringSubmatch(input, -1) {
		min1, _ := strconv.ParseInt(matches[2], 10, 0)
		max1, _ := strconv.ParseInt(matches[3], 10, 0)
		min2, _ := strconv.ParseInt(matches[4], 10, 0)
		max2, _ := strconv.ParseInt(matches[5], 10, 0)
		ranges = append(ranges, Range{matches[1], min1, max1, min2, max2})
	}
	return ranges
}

func getTicketsFromInput(input string) []Ticket {
	tickets := make([]Ticket, 0)
	ticketRe := regexp.MustCompile(`\d+(,\d+)+`)
	for _, match := range ticketRe.FindAllString(input, -1) {
		ticketValues := make(Ticket, 0)
		for _, value := range strings.Split(match, ",") {
			v, _ := strconv.ParseInt(value, 10, 0)
			ticketValues = append(ticketValues, v)
		}
		tickets = append(tickets, ticketValues)
	}
	return tickets
}

func getInput(r io.Reader) ([]Range, []Ticket) {
	b, _ := ioutil.ReadAll(r)
	input := string(b)

	return getRangesFromInput(input), getTicketsFromInput(input)
}

func validForAnyRange(value int64, ranges []Range) bool {
	for _, r := range ranges {
		if r.validate(value) {
			return true
		}
	}
	return false
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	ranges, tickets := getInput(stdinReader)
	positions := makeFieldPositions(ranges)

	var errorRate int64
	for _, ticket := range tickets[1:] {
		validTicket := true
		for _, value := range ticket {
			if !validForAnyRange(value, ranges) {
				errorRate += value
				validTicket = false
			}
		}
		if validTicket {
			positions.updatePossibilities(ticket)
		}
	}
	fmt.Println("The error rate is:", errorRate)

	var product int64
	product = 1
	for i, p := range positions.possibilities {
		for j, _ := range p {
			name := positions.ranges[i].name
			if len(name) >= 9 && name[:10] == "departure " {
				product *= tickets[0][j]
			}
		}
	}
	fmt.Println("The six departure values multiplied:", product)
}
