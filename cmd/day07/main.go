package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagType struct {
	adjective string
	color     string
}

type BagCount struct {
	qty     int
	bagType BagType
}

type Day07Input map[BagType][]BagCount

func getInput(rawBagSpecs []string) Day07Input {
	result := make(map[BagType][]BagCount)
	regexMain := regexp.MustCompile(`(\w+) (\w+) bags contain (.+)$`)
	regexBagSpecs := regexp.MustCompile(`(\d+) (\w+) (\w+) bags?`)
	for _, spec := range rawBagSpecs {
		if len(spec) == 0 {
			continue
		}
		submatches := regexMain.FindStringSubmatch(spec)
		container := BagType{
			adjective: submatches[1],
			color:     submatches[2],
		}
		contents := make([]BagCount, 0)

		if submatches[3] != "no other bags" {
			for _, contentsSpec := range regexBagSpecs.FindAllStringSubmatch(submatches[3], -1) {
				bagType := BagType{
					adjective: contentsSpec[2],
					color:     contentsSpec[3],
				}
				qty, _ := strconv.ParseInt(contentsSpec[1], 10, 0)
				contentsSpec := BagCount{
					qty:     int(qty),
					bagType: bagType,
				}
				contents = append(contents, contentsSpec)
			}
		}

		result[container] = contents
	}
	return result
}

func partOne(input Day07Input, bagType BagType) map[BagType]bool {
	types := make(map[BagType]bool)
	for container, contents := range input {
		for _, count := range contents {
			if count.bagType == bagType {
				types[container] = true
				for t, _ := range partOne(input, container) {
					types[t] = true
				}
			}
		}
	}
	return types
}

func partTwo(input Day07Input, bagType BagType) int {
	qty := 1

	for _, contents := range input[bagType] {
		qty += contents.qty * partTwo(input, contents.bagType)
	}

	return qty
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	rawBagSpecs := strings.Split(string(buffer), ".\n")

	input := getInput(rawBagSpecs)
	shinyGold := BagType{
		adjective: "shiny",
		color:     "gold",
	}

	fmt.Println(len(partOne(input, shinyGold)))
	fmt.Println(partTwo(input, shinyGold) - 1)
}
