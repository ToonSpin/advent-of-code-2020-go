package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Tile [2]int
type Color bool

const East = 0
const Northeast = 1
const Northwest = 2
const West = 3
const Southwest = 4
const Southeast = 5

const Black = true
const White = false

func getInput(r io.Reader) [][]int {
	b, _ := ioutil.ReadAll(r)
	input := make([][]int, 0)

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			break
		}
		tilespec := make([]int, 0)
		for i := 0; i < len(line); i++ {
			if line[i] == 'e' {
				tilespec = append(tilespec, East)
				continue
			}
			if line[i] == 'w' {
				tilespec = append(tilespec, West)
				continue
			}

			if line[i] == 'n' {
				if line[i+1] == 'e' {
					tilespec = append(tilespec, Northeast)
				} else {
					tilespec = append(tilespec, Northwest)
				}
			} else {
				if line[i+1] == 'e' {
					tilespec = append(tilespec, Southeast)
				} else {
					tilespec = append(tilespec, Southwest)
				}
			}
			i++
		}
		input = append(input, tilespec)
	}

	return input
}

func move(p Tile, direction int) Tile {
	switch direction {
	case East:
		p[0]++
	case Northeast:
		p[1]++
	case Northwest:
		p[0]--
		p[1]++
	case West:
		p[0]--
	case Southwest:
		p[1]--
	case Southeast:
		p[0]++
		p[1]--
	}
	return p
}

func getNeighbors(t Tile) []Tile {
	neighbors := make([]Tile, 6)
	for direction := 0; direction < 6; direction++ {
		neighbors[direction] = move(t, direction)
	}
	return neighbors
}

func countBlackNeighbors(t Tile, tiles map[Tile]bool) int {
	count := 0
	for _, n := range getNeighbors(t) {
		value, _ := tiles[n]
		if black(value) {
			count++
		}
	}
	return count
}

func black(value bool) bool {
	return value
}

func white(value bool) bool {
	return !value
}

func iterate(tiles map[Tile]bool) (map[Tile]bool, int) {
	newTiles := make(map[Tile]bool)
	for t, _ := range tiles {
		newTiles[t] = White
		for _, n := range getNeighbors(t) {
			newTiles[n] = White
		}
	}
	count := 0
	for t, _ := range newTiles {
		c := countBlackNeighbors(t, tiles)
		if c != 1 && c != 2 && black(tiles[t]) {
			newTiles[t] = White
		} else if c == 2 && white(tiles[t]) {
			newTiles[t] = Black
			count++
		} else {
			v, _ := tiles[t]
			newTiles[t] = v
			if black(v) {
				count++
			}
		}
	}
	return newTiles, count
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)
	tiles := make(map[Tile]bool)

	for _, path := range input {
		var p Tile
		for _, direction := range path {
			p = move(p, direction)
		}
		f, _ := tiles[p]
		tiles[p] = !f
	}

	count := 0
	for _, f := range tiles {
		if black(f) {
			count++
		}
	}
	fmt.Println("Tiles left with the black side up:", count)

	for i := 1; i <= 100; i++ {
		tiles, count = iterate(tiles)
	}
	fmt.Println("Black tiles after 100 days:", count)
}
