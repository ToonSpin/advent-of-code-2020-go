package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	opcode  string
	operand int
}

func getInput(instructionSpecs string) []Instruction {
	result := make([]Instruction, 0)
	re := regexp.MustCompile(`(acc|jmp|nop) ((?:\+|-)\d+)`)
	for _, instructionSpec := range re.FindAllStringSubmatch(instructionSpecs, -1) {
		operand, _ := strconv.ParseInt(instructionSpec[2], 10, 0)
		i := Instruction{
			opcode:  instructionSpec[1],
			operand: int(operand),
		}
		result = append(result, i)
	}

	return result
}

func runProgram(input []Instruction, partTwo bool) int {
	sp := 0
	acc := 0
	seen := make(map[int]bool)
	partOne := !partTwo
	for {
		_, repeat := seen[sp]
		if partOne && repeat {
			return acc
		}

		seen[sp] = true
		i := input[sp]
		jmpVal := 1

		switch i.opcode {
		case "jmp":
			jmpVal = i.operand
			if partTwo {
				_, repeat := seen[sp+jmpVal]
				if repeat {
					jmpVal = 1
				}
			}
		case "acc":
			acc += i.operand
		case "nop":
			if partTwo && sp+i.operand == len(input) {
				return acc
			}
		}

		sp += jmpVal
		if sp == len(input) {
			return acc
		}
	}
	return 0
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	input := getInput(string(buffer))
	fmt.Println(runProgram(input, false))
	fmt.Println(runProgram(input, true))
}
