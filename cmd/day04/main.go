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

func validatePartOne(passport map[string]string) bool {
	mandatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range mandatoryFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func validatePartTwo(passport map[string]string) bool {
	byr, _ := passport["byr"]
	byrNum, err := strconv.ParseInt(byr, 10, 0)
	if err != nil || byrNum < 1920 || byrNum > 2002 {
		return false
	}

	iyr, _ := passport["iyr"]
	iyrNum, err := strconv.ParseInt(iyr, 10, 0)
	if err != nil || iyrNum < 2010 || iyrNum > 2020 {
		return false
	}

	eyr, _ := passport["eyr"]
	eyrNum, err := strconv.ParseInt(eyr, 10, 0)
	if err != nil || eyrNum < 2020 || eyrNum > 2030 {
		return false
	}

	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	hgt, _ := passport["hgt"]
	matches := re.FindStringSubmatch(hgt)
	if matches == nil {
		return false
	}
	hgtNum, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return false
	}

	if matches[2] == "cm" && (hgtNum < 150 || hgtNum > 193) {
		return false
	}
	if matches[2] == "in" && (hgtNum < 59 || hgtNum > 76) {
		return false
	}

	hcl, _ := passport["hcl"]
	hclMatch, err := regexp.MatchString(`^#[0-9a-z]{6}$`, hcl)
	if err != nil || !hclMatch {
		return false
	}

	ecl, _ := passport["ecl"]
	eclMatch, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, ecl)
	if err != nil || !eclMatch {
		return false
	}

	pid, _ := passport["pid"]
	pidMatch, err := regexp.MatchString(`^\d{9}$`, pid)
	if err != nil || !pidMatch {
		return false
	}

	return true
}

func getInput(r io.Reader) []map[string]string {
	b, _ := ioutil.ReadAll(r)
	rawInput := string(b)
	input := strings.Split(rawInput, "\n\n")

	result := make([]map[string]string, 0)
	re := regexp.MustCompile(`([a-z]{3}):(\S+)`)

	for _, passportString := range input {
		passportFields := make(map[string]string)
		for _, pair := range re.FindAllStringSubmatch(passportString, -1) {
			passportFields[pair[1]] = pair[2]
		}
		result = append(result, passportFields)
	}
	return result
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)

	numValidPartOne := 0
	numValidPartTwo := 0

	for _, passport := range input {
		if validatePartOne(passport) {
			numValidPartOne++
			if validatePartTwo(passport) {
				numValidPartTwo++
			}
		}
	}

	fmt.Println("Number of passports with all mandatory fields:", numValidPartOne)
	fmt.Println("Number of passports with valid field values:", numValidPartTwo)
}
