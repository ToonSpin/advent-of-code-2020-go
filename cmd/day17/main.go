package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Coords struct {
	x int
	y int
	z int
	w int
}

type CubeCluster map[Coords]bool

func (d CubeCluster) iterate(partTwo bool) CubeCluster {
	new := make(CubeCluster)
	neighborhood := make(CubeCluster)

	for p, _ := range d {
		numActive := 0
		for _, q := range getNeighbors(p, partTwo) {
			neighborhood[q] = true
			_, active := d[q]
			if active {
				numActive++
			}
		}
		if numActive == 2 || numActive == 3 {
			new[p] = true
		}
	}

	for p, _ := range neighborhood {
		_, active := d[p]
		if active {
			continue
		}

		numActive := 0
		for _, q := range getNeighbors(p, partTwo) {
			_, active := d[q]
			if active {
				numActive++
			}
		}
		if numActive == 3 {
			new[p] = true
		}
	}

	return new
}

func getNeighbors(c Coords, partTwo bool) []Coords {
	result := make([]Coords, 0)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				var t int
				if partTwo {
					t = 1
				}

				for w := c.w - t; w <= c.w+t; w++ {
					if x == c.x && y == c.y && z == c.z && w == c.w {
						continue
					}
					result = append(result, Coords{x, y, z, w})
				}
			}
		}
	}
	return result
}

func bootstrap(pocketDimension CubeCluster, partTwo bool) int {
	for i := 0; i < 6; i++ {
		pocketDimension = pocketDimension.iterate(partTwo)
	}
	return len(pocketDimension)
}

func getInput(rawInput string) CubeCluster {
	pocketDimension := make(CubeCluster)
	for y, line := range strings.Split(rawInput, "\n") {
		if len(line) == 0 {
			continue
		}
		for x, r := range line {
			if r == '#' {
				coords := Coords{x, y, 0, 0}
				pocketDimension[coords] = true
			}
		}
	}
	return pocketDimension
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	pocketDimension := getInput(string(buffer))

	fmt.Println("The number of cubes in three dimensions:", bootstrap(pocketDimension, false))
	fmt.Println("The number of cubes in four dimensions:", bootstrap(pocketDimension, true))
}
