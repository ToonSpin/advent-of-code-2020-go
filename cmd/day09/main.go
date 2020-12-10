package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const PreambleSize = 25

func getInput() []int64 {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	input := make([]int64, 0, 1000)

	for _, s := range strings.Split(string(buffer), "\n") {
		if len(s) > 0 {
			n, _ := strconv.ParseInt(s, 10, 0)
			input = append(input, n)
		}
	}

	return input
}

func partOne(input []int64) (int64, error) {
	for currentIndex := PreambleSize; currentIndex < len(input); currentIndex++ {
		curSl := input[currentIndex-PreambleSize : currentIndex]
		candidate := input[currentIndex]
		found := false

		for i := 0; i < PreambleSize && !found; i++ {
			for j := 0; j < i && !found; j++ {
				if curSl[i]+curSl[j] == candidate {
					found = true
				}
			}
		}

		if !found {
			return candidate, nil
		}
	}
	return 0, errors.New("No invalid number found")
}

func minmax(input []int64) (int64, int64) {
	min := input[0]
	max := min

	for _, n := range input {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}

func partTwo(input []int64, invalid int64, span int) (int64, int64, error) {
	var sum int64
	for i := 0; i < span; i++ {
		sum += input[i]
	}

	for i := span; i < len(input); i++ {
		// putting this block here means I have to assume that the answer won't span the entire input
		if sum == invalid {
			min, max := minmax(input[i-span : i])
			return min, max, nil
		}
		sum -= input[i-span]
		sum += input[i]
	}

	return 0, 0, errors.New("No encryption weakness found")
}

func main() {
	input := getInput()

	invalid, _ := partOne(input)
	fmt.Println("The invalid number is:", invalid)

	for span := 2; span < len(input); span++ {
		lower, upper, err := partTwo(input, invalid, span)

		if err == nil {
			fmt.Println("The encryption weakness is:", lower+upper)
			break
		}
	}
}
