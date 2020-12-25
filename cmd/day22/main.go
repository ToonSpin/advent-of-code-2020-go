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

func getInput(r io.Reader) [][]int {
	b, _ := ioutil.ReadAll(r)
	input := string(b)
	decks := make([][]int, 0)

	deckRe := regexp.MustCompile(`(\d+\n)+`)
	for _, deckStr := range deckRe.FindAllString(input, -1) {
		deck := make([]int, 0)
		for _, card := range strings.Split(deckStr, "\n") {
			if len(card) == 0 {
				continue
			}
			card, _ := strconv.ParseInt(card, 10, 0)
			deck = append(deck, int(card))
		}
		decks = append(decks, deck)
	}

	return decks
}

func duplicateDecks(decks [][]int) [][]int {
	decks2 := make([][]int, len(decks))
	for i, r := range decks {
		decks2[i] = make([]int, len(r))
		copy(decks2[i], r)
	}
	return decks2
}

func getScore(decks [][]int, winningDeck int) int {
	score := 0
	numInDeck := len(decks[winningDeck])
	for i, card := range decks[winningDeck] {
		score += (numInDeck - i) * card
	}
	return score
}

func playGame(decks [][]int, partTwo bool) (winningPlayer, score int) {
	statesSeen := make(map[string]bool)
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		if partTwo {
			state := fmt.Sprintf("%v", decks)
			_, seen := statesSeen[state]
			if seen {
				winningPlayer = 0
				score = getScore(decks, winningPlayer)
				return
			}
			statesSeen[state] = true
		}

		card0 := decks[0][0]
		card1 := decks[1][0]
		decks[0] = decks[0][1:]
		decks[1] = decks[1][1:]

		var winner int
		if partTwo && card0 <= len(decks[0]) && card1 <= len(decks[1]) {
			subDecks := duplicateDecks(decks)
			subDecks[0] = subDecks[0][:card0]
			subDecks[1] = subDecks[1][:card1]
			winner, _ = playGame(subDecks, partTwo)
		} else {
			if card0 > card1 {
				winner = 0
			} else {
				winner = 1
			}
		}

		if winner == 0 {
			decks[0] = append(decks[0], card0)
			decks[0] = append(decks[0], card1)
		} else {
			decks[1] = append(decks[1], card1)
			decks[1] = append(decks[1], card0)
		}
	}

	var winningDeck int
	if len(decks[0]) > 0 {
		winningDeck = 0
	} else {
		winningDeck = 1
	}

	return winningDeck, getScore(decks, winningDeck)
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)

	_, scorePartOne := playGame(duplicateDecks(input), false)
	fmt.Println("Score of winner of Combat game:", scorePartOne)
	_, scorePartTwo := playGame(duplicateDecks(input), true)
	fmt.Println("Score of winner of Recursive Combat game:", scorePartTwo)
}
