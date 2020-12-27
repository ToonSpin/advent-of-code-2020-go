package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TileId int64
type RowCol [10]bool
type Grid [10][10]bool

func (g Grid) matchHorizontal(h Grid) bool {
	for y := 0; y < 10; y++ {
		if g[y][9] != h[y][0] {
			return false
		}
	}
	return true
}

func (g Grid) matchVertical(h Grid) bool {
	return g[9] == h[0]
}

func (g Grid) String() string {
	result := ""
	for _, row := range g {
		for _, b := range row {
			if b {
				result = fmt.Sprintf("%s#", result)
			} else {
				result = fmt.Sprintf("%s.", result)
			}
		}
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

type TilePermutation struct {
	id     TileId
	flip   bool
	rotate int
	grid   Grid
}

func (candidate TilePermutation) tileFits(partial []TilePermutation, rowSize int) bool {
	row := len(partial) / rowSize
	col := len(partial) % rowSize

	if row > 0 {
		aboveIndex := (row-1)*rowSize + col
		if !partial[aboveIndex].grid.matchVertical(candidate.grid) {
			return false
		}
	}

	if col > 0 {
		leftIndex := len(partial) - 1
		if !partial[leftIndex].grid.matchHorizontal(candidate.grid) {
			return false
		}
	}

	return true
}

type Tile struct {
	id   TileId
	grid Grid
}

func (t Tile) getPermutations() []TilePermutation {
	result := make([]TilePermutation, 0)
	for rotation := 0; rotation <= 3; rotation++ {
		result = append(result, t.makePermutation(false, rotation))
		result = append(result, t.makePermutation(true, rotation))
	}
	return result
}

func (t Tile) makePermutation(flip bool, rotate int) TilePermutation {
	grid := t.grid
	if flip {
		var newGrid Grid
		for y := 0; y < 10; y++ {
			newGrid[9-y] = grid[y]
		}
		grid = newGrid
	}
	for i := 0; i < rotate; i++ {
		var newGrid Grid
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				p := 9 - y
				q := x
				newGrid[q][p] = grid[y][x]
			}
		}
		grid = newGrid
	}
	return TilePermutation{t.id, flip, rotate, grid}
}

type Image struct {
	rowSize int
	image   [][]bool
}

func (im *Image) flip() {
	newImage := make([][]bool, 0)
	for r := 0; r < im.rowSize; r++ {
		row := make([]bool, im.rowSize)
		copy(row, im.image[im.rowSize-r-1])
		newImage = append(newImage, row)
	}
	im.image = newImage
}

func (im *Image) rotate() {
	newImage := make([][]bool, 0)
	for r := 0; r < im.rowSize; r++ {
		row := make([]bool, 0, im.rowSize)
		for c := 0; c < im.rowSize; c++ {
			rr := im.rowSize - c - 1
			cc := r
			row = append(row, im.image[rr][cc])
		}
		newImage = append(newImage, row)
	}
	im.image = newImage
}

func (im Image) matchSeaMonster(rowIndex, columnIndex int) bool {
	seaMonster := [][]int{{18}, {0, 5, 6, 11, 12, 17, 18, 19}, {1, 4, 7, 10, 13, 16}}
	seaMonsterWidth := 20
	seaMonsterHeight := 3

	if rowIndex+seaMonsterHeight > im.rowSize {
		return false
	}
	if columnIndex+seaMonsterWidth > im.rowSize {
		return false
	}

	for r, positions := range seaMonster {
		for _, c := range positions {
			if !im.image[rowIndex+r][columnIndex+c] {
				return false
			}
		}
	}

	return true
}

func tileIdPresent(id TileId, perms []TilePermutation) bool {
	for _, p := range perms {
		if p.id == id {
			return true
		}
	}
	return false
}

func arrangeTiles(partial []TilePermutation, rowSize int, input []Tile) ([]TilePermutation, error) {
	if len(partial) == rowSize*rowSize {
		return partial, nil
	}

	for _, candidateTile := range input {
		if tileIdPresent(candidateTile.id, partial) {
			continue
		}

		for _, perm := range candidateTile.getPermutations() {
			if perm.tileFits(partial, rowSize) {
				partial = append(partial, perm)
				result, err := arrangeTiles(partial, rowSize, input)
				if err == nil {
					return result, nil
				}
				partial = partial[:len(partial)-1]
			}
		}
	}

	return partial, errors.New("Could not find a fitting arrangement")
}

func getImage(arrangement []TilePermutation) Image {
	image := make([][]bool, 0)
	var rowSize int
	for rowSize = 0; rowSize*rowSize != len(arrangement); rowSize++ {
	}

	for r := 0; r < rowSize; r++ {
		for line := 1; line <= 8; line++ {
			imageRow := make([]bool, 0, rowSize*10)
			for c := 0; c < rowSize; c++ {
				tile := arrangement[r*rowSize+c]
				imageRow = append(imageRow, tile.grid[line][1:9]...)
			}
			image = append(image, imageRow)
		}
	}
	return Image{len(image[0]), image}
}

func getPartOne(arrangement []TilePermutation, rowSize int) TileId {
	var product TileId = 1
	bottomRow := (rowSize - 1) * rowSize
	product *= arrangement[0].id
	product *= arrangement[rowSize-1].id
	product *= arrangement[bottomRow].id
	product *= arrangement[bottomRow+rowSize-1].id
	return product
}

func getPartTwo(im Image) int {
	trueCount := 0
	for _, row := range im.image {
		for _, b := range row {
			if b {
				trueCount++
			}
		}
	}
	numMonstersFound := 0
	for r := 0; r < im.rowSize; r++ {
		for c := 0; c < im.rowSize; c++ {
			if im.matchSeaMonster(r, c) {
				numMonstersFound++
			}
		}
	}
	return trueCount - 15*numMonstersFound
}

func getInput(r io.Reader) []Tile {
	b, _ := ioutil.ReadAll(r)
	input := string(b)

	tiles := make([]Tile, 0)

	tileRe := regexp.MustCompile(`Tile (\d+):\n(([#.]{10}\n){10})`)
	for _, matches := range tileRe.FindAllStringSubmatch(input, -1) {
		id, _ := strconv.ParseInt(matches[1], 10, 0)
		var grid Grid
		for y, line := range strings.Split(matches[2], "\n")[:10] {
			for x, r := range line {
				if r == '#' {
					grid[y][x] = true
				} else {
					grid[y][x] = false
				}
			}
		}
		tiles = append(tiles, Tile{TileId(id), grid})
	}

	return tiles
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	input := getInput(stdinReader)
	var rowSize int
	for rowSize = 0; rowSize*rowSize != len(input); rowSize++ {
	}

	arrangement, _ := arrangeTiles(make([]TilePermutation, 0), rowSize, input)
	fmt.Println("The product of the corner tiles:", getPartOne(arrangement, rowSize))

	image := getImage(arrangement)
	min := 10000
	for i := 0; i <= 7; i++ {
		if i%4 == 0 {
			image.flip()
		}
		image.rotate()

		count := getPartTwo(image)
		if min > count {
			min = count
		}
	}
	fmt.Println("The habitat's water roughness:", min)
}
