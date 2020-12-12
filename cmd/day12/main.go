package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

const North = 0
const East = 1
const South = 2
const West = 3

type Instruction struct {
	instr string
	param int
}

func getInput(instructionSpecs string) []Instruction {
	result := make([]Instruction, 0)
	re := regexp.MustCompile(`(\w)(\d+)`)

	for _, instructionSpec := range re.FindAllStringSubmatch(instructionSpecs, -1) {
		param, _ := strconv.ParseInt(instructionSpec[2], 10, 0)
		i := Instruction{
			instr: instructionSpec[1],
			param: int(param),
		}
		result = append(result, i)
	}

	return result
}

func goNorth(x, y, n int) (x2, y2 int) {
	x2 = x
	y2 = y + n
	return
}

func goEast(x, y, n int) (x2, y2 int) {
	x2 = x + n
	y2 = y
	return
}

func goSouth(x, y, n int) (x2, y2 int) {
	x2 = x
	y2 = y - n
	return
}

func goWest(x, y, n int) (x2, y2 int) {
	x2 = x - n
	y2 = y
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func followInstPartOne(input []Instruction, initialDir int) (int, int) {
	var x, y int
	dir := initialDir
	for _, instruction := range input {
		switch instruction.instr {
		case "N":
			x, y = goNorth(x, y, instruction.param)
		case "E":
			x, y = goEast(x, y, instruction.param)
		case "S":
			x, y = goSouth(x, y, instruction.param)
		case "W":
			x, y = goWest(x, y, instruction.param)
		case "F":
			switch dir {
			case North:
				x, y = goNorth(x, y, instruction.param)
			case East:
				x, y = goEast(x, y, instruction.param)
			case West:
				x, y = goWest(x, y, instruction.param)
			case South:
				x, y = goSouth(x, y, instruction.param)
			}
		case "L":
			p := instruction.param / 90
			dir = dir - p + 16
			dir %= 4
		case "R":
			p := instruction.param / 90
			dir += p
			dir %= 4
		}
	}
	return x, y
}

func followInstPartTwo(input []Instruction, wpx, wpy int) (int, int) {
	var x, y int
	for _, instruction := range input {
		switch instruction.instr {
		case "N":
			wpx, wpy = goNorth(wpx, wpy, instruction.param)
		case "E":
			wpx, wpy = goEast(wpx, wpy, instruction.param)
		case "S":
			wpx, wpy = goSouth(wpx, wpy, instruction.param)
		case "W":
			wpx, wpy = goWest(wpx, wpy, instruction.param)
		case "F":
			x += instruction.param * wpx
			y += instruction.param * wpy
		case "L":
			a := instruction.param / 90
			for i := 0; i < a; i++ {
				temp := wpx
				wpx = -wpy
				wpy = temp
			}
		case "R":
			a := instruction.param / 90
			for i := 0; i < a; i++ {
				temp := wpx
				wpx = wpy
				wpy = -temp
			}
		}
	}
	return x, y
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	input := getInput(string(buffer))

	x, y := followInstPartOne(input, East)
	fmt.Println("Distance according to guessed instructions:", abs(x)+abs(y))

	x, y = followInstPartTwo(input, 10, 1)
	fmt.Println("Distance according to actual instructions:", abs(x)+abs(y))
}
