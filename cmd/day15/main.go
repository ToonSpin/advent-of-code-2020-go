package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type TurnInfo map[int64][2]int64

func (ti TurnInfo) getTurns(i int64) [2]int64 {
	val, found := ti[i]
	if !found {
		val = [2]int64{-1, -1}
		ti[i] = val
	}
	return val
}

func (ti TurnInfo) getTurnDiff(i int64) int64 {
	turns := ti.getTurns(i)
	if turns[0] == -1 && turns[1] == -1 {
		return 0
	}
	if turns[0] == -1 || turns[1] == -1 {
		return 0
	}
	return turns[0] - turns[1]
}

func (ti TurnInfo) speak(i int64, turn int64) {
	turns := ti.getTurns(i)
	turns[1] = turns[0]
	turns[0] = turn
	ti[i] = turns
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)

	input := make([]int64, 0)
	firstLine := strings.Split(string(buffer), "\n")[0]
	for _, s := range strings.Split(firstLine, ",") {
		i, _ := strconv.ParseInt(s, 10, 0)
		input = append(input, i)
	}

	turnInfo := make(TurnInfo)
	for turnIndex, number := range input {
		turnInfo.speak(number, int64(turnIndex+1))
	}

	var turn int64
	prevSpoken := input[len(input)-1]
	for turn = int64(len(input)) + 1; turn <= 30000000; turn++ {
		nextSpoken := turnInfo.getTurnDiff(prevSpoken)
		turnInfo.speak(nextSpoken, turn)
		prevSpoken = nextSpoken

		if turn == 2020 {
			fmt.Println("The number spoken on turn 2020:", prevSpoken)
		}
	}
	fmt.Println("The number spoken on turn 30000000:", prevSpoken)
}
