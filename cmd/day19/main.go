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

type Message string

type Rule interface {
	Match(s string, rules map[int64]Rule) (bool, string)
}

type Literal struct {
	literal string
}

func (r Literal) Match(s string, rules map[int64]Rule) (bool, string) {
	if len(s) < len(r.literal) {
		return false, s
	}
	matchingPart := s[:len(r.literal)]
	if matchingPart != r.literal {
		return false, s
	}
	return true, s[len(r.literal):]
}

type Compound struct {
	subRules []int64
}

func (r Compound) Match(s string, rules map[int64]Rule) (bool, string) {
	rest := s
	matched := true
	for _, subRule := range r.subRules {
		rule := rules[subRule]
		matched, rest = rule.Match(rest, rules)
		if !matched {
			return false, s
		}
	}
	return true, rest
}

type Options struct {
	subRules []int64
}

func (r Options) Match(s string, rules map[int64]Rule) (bool, string) {
	for _, subRule := range r.subRules {
		matched, rest := rules[subRule].Match(s, rules)
		if matched {
			return true, rest
		}
	}
	return false, s
}

func registerCompoundRule(rules map[int64]Rule, subRules string, id int64) {
	ruleIds := make([]int64, 0)
	for _, s := range strings.Split(subRules, " ") {
		subRuleId, _ := strconv.ParseInt(s, 10, 0)
		ruleIds = append(ruleIds, subRuleId)
	}
	rules[id] = Compound{ruleIds}
}

func getInput(r io.Reader) ([]Message, map[int64]Rule) {
	b, _ := ioutil.ReadAll(r)
	input := string(b)

	messages := make([]Message, 0)
	rules := make(map[int64]Rule)

	msgRe := regexp.MustCompile(`[ab]{2,}`)
	for _, match := range msgRe.FindAllString(input, -1) {
		messages = append(messages, Message(match))
	}

	litRe := regexp.MustCompile(`(\d+): "([ab])"`)
	for _, matches := range litRe.FindAllStringSubmatch(input, -1) {
		ruleId, _ := strconv.ParseInt(matches[1], 10, 0)
		rules[ruleId] = Literal{matches[2]}
	}

	cpdRe := regexp.MustCompile(`(?m)^(\d+): (\d+(?: \d+)*)+$`)
	for _, matches := range cpdRe.FindAllStringSubmatch(input, -1) {
		ruleId, _ := strconv.ParseInt(matches[1], 10, 0)
		registerCompoundRule(rules, matches[2], ruleId)
	}

	dummyRuleId := int64(-1)
	optRe := regexp.MustCompile(`(?m)^(\d+): (\d+(?: \d+)*) \| (\d+(?: \d+)*)$`)
	for _, matches := range optRe.FindAllStringSubmatch(input, -1) {
		ruleId, _ := strconv.ParseInt(matches[1], 10, 0)

		registerCompoundRule(rules, matches[2], dummyRuleId)
		registerCompoundRule(rules, matches[3], dummyRuleId-1)

		rules[ruleId] = Options{[]int64{dummyRuleId, dummyRuleId - 1}}
		dummyRuleId -= 2
	}

	return messages, rules
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)

	messages, rules := getInput(stdinReader)
	count := 0
	for _, m := range messages {
		matched, rest := rules[0].Match(string(m), rules)
		if matched && len(rest) == 0 {
			count++
		}
	}
	fmt.Println("Number of valid messages:", count)

	count = 0
	for _, m := range messages {
		rest := string(m)
		var matched bool

		count42 := -1
		matched = true
		for matched {
			matched, rest = rules[42].Match(rest, rules)
			count42++
		}

		count31 := -1
		matched = true
		for matched {
			matched, rest = rules[31].Match(rest, rules)
			count31++
		}

		if len(rest) > 0 {
			continue
		}

		if count42 > count31 && count31 != 0 {
			count++
		}
	}
	fmt.Println("Number of valid messages after rule substitution:", count)
}
