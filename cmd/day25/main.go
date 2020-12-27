package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func transform(subject, loopSize int64) int64 {
	value := int64(1)
	for i := int64(0); i < loopSize; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

func findLoopSize(subject, key int64) int64 {
	value := int64(1)
	i := int64(0)
	for ; value != key; i++ {
		value *= subject
		value %= 20201227
	}
	return i
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(stdinReader)
	rawInput := strings.Split(string(b), "\n")

	key1, _ := strconv.ParseInt(rawInput[0], 10, 0)
	key2, _ := strconv.ParseInt(rawInput[1], 10, 0)

	loopsize1 := findLoopSize(7, key1)
	fmt.Println("The encryption key is:", transform(key2, loopsize1))
}
