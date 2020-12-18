package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Expression interface {
	compute(partTwo bool) int64
}

type Constant struct {
	value int64
}

func (c Constant) compute(partTwo bool) int64 {
	return c.value
}

type Compound struct {
	operators []byte
	operands  []Expression
}

func (c Compound) compute(partTwo bool) int64 {
	if partTwo {
		return c.computePartTwo()
	}
	return c.computePartOne()
}

func (c Compound) computePartOne() int64 {
	accumulator := c.operands[0].compute(false)
	for i, operator := range c.operators {
		switch operator {
		case '+':
			accumulator += c.operands[i+1].compute(false)
		case '*':
			accumulator *= c.operands[i+1].compute(false)
		default:
			panic("Unknown operator")
		}
	}
	return accumulator
}

func (c Compound) computePartTwo() int64 {
	mulOperands := make([]int64, 0)
	operand := c.operands[0].compute(true)
	for i, operator := range c.operators {
		value := c.operands[i+1].compute(true)
		switch operator {
		case '+':
			operand += value
		case '*':
			mulOperands = append(mulOperands, operand)
			operand = value
		default:
			panic("Unknown operator")
		}
	}
	mulOperands = append(mulOperands, operand)

	var accumulator int64
	accumulator = 1
	for _, o := range mulOperands {
		accumulator *= o
	}

	return accumulator
}

func parseWs(s string) string {
	var j int
	for ; j < len(s); j++ {
		if s[j] != ' ' {
			break
		}
	}
	return s[j:]
}

func parseConstant(s string) (Constant, string) {
	var j int
	for ; j < len(s); j++ {
		if s[j] < '0' || s[j] > '9' {
			break
		}
	}
	n, _ := strconv.ParseInt(s[:j], 10, 0)
	return Constant{n}, s[j:]
}

func firstUnmatchedParenthesis(s string) int {
	count := 0
	for i, r := range s {
		if r == ')' {
			if count == 0 {
				return i
			}
			count--
		}
		if r == '(' {
			count++
		}
	}
	panic("no unmatched parenthesis")
	return 0
}

func parseOperand(s string) (Expression, string) {
	var e Expression
	if s[0] == '(' {
		matchingParen := firstUnmatchedParenthesis(s[1:]) + 1
		temp, _ := parseExpression(s[1:matchingParen])
		e = temp
		s = s[matchingParen+1:]
	} else if s[0] >= '0' && s[0] <= '9' {
		e, s = parseConstant(s)
	} else {
		panic("can't parse this")
	}
	s = parseWs(s)
	return e, s
}

func parseExpression(s string) (Expression, string) {
	var e Expression
	e, s = parseOperand(s)
	if s == "" {
		return e, s
	}

	operators := make([]byte, 0)
	operands := make([]Expression, 1)
	operands[0] = e

	for len(s) > 0 {
		var operand Expression
		operators = append(operators, s[0])
		s = parseWs(s[1:])
		operand, s = parseOperand(s)
		operands = append(operands, operand)
		s = parseWs(s)
	}
	return Compound{operators, operands}, s
}

func getInput(rawInput string) []Expression {
	expressions := make([]Expression, 0)
	for _, line := range strings.Split(rawInput, "\n") {
		if len(line) == 0 {
			continue
		}
		e, _ := parseExpression(line)
		expressions = append(expressions, e)
	}
	return expressions
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	input := getInput(string(buffer))

	var sumPartOne int64
	var sumPartTwo int64
	for _, e := range input {
		sumPartOne += e.compute(false)
		sumPartTwo += e.compute(true)
	}

	fmt.Println("The regular math sum of the expressions:", sumPartOne)
	fmt.Println("The advanced math sum of the expressions:", sumPartTwo)
}
