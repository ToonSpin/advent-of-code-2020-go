package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const maxLabelPartTwo = 1000000
const numMovesPartTwo = 10000000

type Cup struct {
	label int
	prev  int
	next  int
}

type Circle struct {
	circle      []Cup
	current     int
	minLabel    int
	maxLabel    int
	cupsByLabel map[int]int
}

func (c Circle) getCurrentLabel() int {
	return c.circle[c.current].label
}

func (c Circle) getNextLabelForIteration(pickupIndices [3]int) int {
	nextLabel := c.circle[c.current].label - 1
	for done := false; !done; {
		done = true
		if nextLabel < c.minLabel {
			nextLabel = c.maxLabel
		}
		for _, index := range pickupIndices {
			if c.circle[index].label == nextLabel {
				done = false
				nextLabel--
				break
			}
		}
	}
	return nextLabel
}

func (c Circle) print() {
	start := c.current
	for ; c.circle[start].label != 1; start = c.circle[start].next {
	}
	fmt.Print("cups: ")
	cursor := start
	for {
		if cursor == c.current {
			fmt.Printf("(%d)", c.circle[cursor].label)
		} else {
			fmt.Printf(" %d ", c.circle[cursor].label)
		}
		cursor = c.circle[cursor].next
		if cursor == start {
			break
		}
	}
	fmt.Println()
}

func (c Circle) printPartOne() {
	cursor := 0
	for ; c.circle[cursor].label != 1; cursor = c.circle[cursor].next {
	}
	cursor = c.circle[cursor].next
	for {
		if c.circle[cursor].label == 1 {
			break
		}
		fmt.Printf("%d", c.circle[cursor].label)
		cursor = c.circle[cursor].next
	}
	fmt.Println()
}

func (c Circle) printPartTwo() {
	cursor := c.cupsByLabel[1]
	cursor = c.circle[cursor].next

	product := c.circle[cursor].label
	cursor = c.circle[cursor].next
	product *= c.circle[cursor].label
	fmt.Println(product)
}

func (c *Circle) iterate() {
	var pickupIndices [3]int
	cursor := c.current
	for i := 0; i <= 2; i++ {
		cursor = c.circle[cursor].next
		pickupIndices[i] = cursor
	}

	c.circle[c.current].next = c.circle[cursor].next
	c.circle[c.circle[c.current].next].prev = c.current

	nextLabel := c.getNextLabelForIteration(pickupIndices)
	cursor = c.cupsByLabel[nextLabel]

	indexAfter := c.circle[cursor].next
	c.circle[cursor].next = pickupIndices[0]
	c.circle[pickupIndices[0]].prev = cursor

	c.circle[indexAfter].prev = pickupIndices[2]
	c.circle[pickupIndices[2]].next = indexAfter

	c.current = c.circle[c.current].next
}

func getCircle(input []int, partTwo bool) Circle {
	circle := make([]Cup, 0, maxLabelPartTwo)
	cupsByLabel := make(map[int]int)

	var min int
	var max int
	min = maxLabelPartTwo
	for i, label := range input {
		if label > max {
			max = label
		}
		if label < min {
			min = label
		}
		cup := Cup{label, i - 1, i + 1}
		circle = append(circle, cup)
		cupsByLabel[label] = i
	}
	if partTwo {
		i := len(circle)
		for label := max + 1; label <= maxLabelPartTwo; label++ {
			cup := Cup{label, i - 1, i + 1}
			circle = append(circle, cup)
			cupsByLabel[label] = i
			i++
		}
	}
	if partTwo {
		max = maxLabelPartTwo
	}
	circle[len(circle)-1].next = 0
	circle[0].prev = len(circle) - 1
	return Circle{circle, 0, min, max, cupsByLabel}
}

func getInput(r io.Reader) []int {
	b, _ := ioutil.ReadAll(r)
	input := make([]int, 0)

	for _, i := range b {
		if i == byte('\n') {
			break
		}
		input = append(input, int(i-'0'))
	}

	return input
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)

	circle1 := getCircle(input, false)
	for move := 0; move < 100; move++ {
		circle1.iterate()
	}
	fmt.Print("The labels after cup 1: ")
	circle1.printPartOne()

	circle2 := getCircle(input, true)
	for move := 0; move < numMovesPartTwo; move++ {
		circle2.iterate()
	}
	fmt.Print("The product of the two labels after cup 1: ")
	circle2.printPartTwo()
}
