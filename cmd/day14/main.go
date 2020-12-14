package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type State struct {
	mask   string
	memory map[int]string
}

type Instruction interface {
	Execute(state *State, partTwo bool)
}

type SetMaskInstruction struct {
	mask string
}

func (instr SetMaskInstruction) Execute(state *State, partTwo bool) {
	state.mask = instr.mask
}

type SetMemInstruction struct {
	address int
	value   string
}

func (instr SetMemInstruction) Execute(state *State, partTwo bool) {
	if partTwo {
		address := intToString(instr.address)
		applied := applyMaskToAddress(address, state.mask)
		acc := make([]int, 0)
		getMemoryAddresses(applied, &acc)
		for _, addr := range acc {
			state.memory[addr] = instr.value
		}
	} else {
		newValue := applyMaskToValue(instr.value, state.mask)
		state.memory[instr.address] = newValue
	}
}

func intToString(n int) string {
	return fmt.Sprintf("%036b", n)
}

func stringToInt(n string) int {
	return int(stringToInt64(n))
}

func stringToInt64(n string) int64 {
	i, _ := strconv.ParseInt(n, 2, 0)
	return i
}

func getMemoryAddresses(mask string, acc *[]int) {
	bytes := []byte(mask)
	for i, r := range []byte(bytes) {
		if r == 'X' {
			bytes[i] = '0'
			getMemoryAddresses(string(bytes), acc)
			bytes[i] = '1'
			getMemoryAddresses(string(bytes), acc)
			return
		}
	}
	*acc = append(*acc, stringToInt(mask))
}

func applyMaskToAddress(address, mask string) string {
	result := make([]byte, 36)
	for i, r := range []byte(address) {
		switch mask[i] {
		case '0':
			result[i] = r
		case '1':
			result[i] = '1'
		case 'X':
			result[i] = 'X'
		default:
			panic("Invalid mask bit")
		}
	}
	return string(result)
}

func applyMaskToValue(value, mask string) string {
	result := make([]byte, 36)
	for i, r := range []byte(value) {
		switch mask[i] {
		case '0':
			result[i] = '0'
		case '1':
			result[i] = '1'
		case 'X':
			result[i] = r
		default:
			panic("Invalid mask bit")
		}
	}
	return string(result)
}

func getInput(instructionSpecs string) []Instruction {
	re := regexp.MustCompile(`(mask.*|mem.*)\n`)

	maskRe := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	subMatches := re.FindAllStringSubmatch(instructionSpecs, -1)
	result := make([]Instruction, len(subMatches))

	for i, instructionSpec := range subMatches {
		if matches := memRe.FindStringSubmatch(instructionSpec[1]); matches != nil {
			address, _ := strconv.ParseInt(matches[1], 10, 0)
			v, _ := strconv.ParseInt(matches[2], 10, 0)
			value := intToString(int(v))
			result[i] = SetMemInstruction{int(address), value}
		} else if matches := maskRe.FindStringSubmatch(instructionSpec[1]); matches != nil {
			result[i] = SetMaskInstruction{matches[1]}
		} else {
			panic("Unknown instruction")
		}
	}

	return result
}

func sumOfMemoryValues(input []Instruction, partTwo bool) int64 {
	state := State{"", make(map[int]string)}
	for _, instr := range input {
		instr.Execute(&state, partTwo)
	}

	var sum int64
	for _, value := range state.memory {
		sum += stringToInt64(value)
	}

	return sum
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	input := getInput(string(buffer))

	partOne := sumOfMemoryValues(input, false)
	fmt.Println("Sum of all values in memory (part 1):", partOne)

	partTwo := sumOfMemoryValues(input, true)
	fmt.Println("Sum of all values in memory (part 2):", partTwo)
}
