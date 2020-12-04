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

type Password struct {
	param1    int64
	param2    int64
	character byte
	password  string
}

func makePasswordFromStrArray(a []string) Password {
	param1, _ := strconv.ParseInt(a[1], 10, 0)
	param2, _ := strconv.ParseInt(a[2], 10, 0)
	return Password{
		param1,
		param2,
		a[3][0],
		a[4],
	}
}

func partOne(input []Password) int {
	validPasswordCount := 0

	for _, p := range input {
		count := int64(strings.Count(p.password, string(p.character)))

		if p.param1 <= count && p.param2 >= count {
			validPasswordCount++
		}
	}

	return validPasswordCount
}

func partTwo(input []Password) int {
	validPasswordCount := 0

	for _, p := range input {
		if p.param1-1 > int64(len(p.password)) {
			continue
		}
		if p.param2-1 > int64(len(p.password)) {
			continue
		}
		match1 := p.password[p.param1-1] == p.character
		match2 := p.password[p.param2-1] == p.character

		if match1 && match2 {
			continue
		}
		if match1 || match2 {
			validPasswordCount++
		}
	}

	return validPasswordCount
}

func getInput(r io.Reader) []Password {
	b, _ := ioutil.ReadAll(r)
	input := string(b)
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	rawInput := re.FindAllStringSubmatch(input, -1)
	result := make([]Password, 0)
	for _, p := range rawInput {
		result = append(result, makePasswordFromStrArray(p))
	}
	return result
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)

	var validPasswordCount int

	validPasswordCount = partOne(input)
	fmt.Printf("Found %d valid O.T.C.A.S. passwords\n", validPasswordCount)

	validPasswordCount = partTwo(input)
	fmt.Printf("Found %d valid O.T.C.P. passwords\n", validPasswordCount)
}
