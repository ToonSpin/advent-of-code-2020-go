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

type Bus struct {
	pos int64
	id  int64
}

func getInput() (int64, []Bus) {
	buses := make([]Bus, 0)

	stdinReader := bufio.NewReader(os.Stdin)
	buffer, _ := ioutil.ReadAll(stdinReader)
	rawInput := strings.Split(string(buffer), "\n")
	earliestDeparture, _ := strconv.ParseInt(rawInput[0], 10, 0)

	for i, r := range strings.Split(rawInput[1], ",") {
		if r == "x" {
			continue
		}

		b, _ := strconv.ParseInt(string(r), 10, 0)
		buses = append(buses, Bus{int64(i), b})
	}

	return int64(earliestDeparture), buses
}

func getMaxBus(buses []Bus) Bus {
	maxBus := buses[0]
	for _, bus := range buses[1:] {
		if maxBus.id < bus.id {
			maxBus = bus
		}
	}
	return maxBus
}

func excludeBus(buses []Bus, toExclude Bus) []Bus {
	result := make([]Bus, 0)
	for _, b := range buses {
		if b == toExclude {
			continue
		}
		result = append(result, b)
	}
	return result
}

func partTwo(buses []Bus) (int64, error) {
	maxBus := getMaxBus(buses)
	timestamp := maxBus.id - maxBus.pos
	interval := maxBus.id
	buses = excludeBus(buses, maxBus)

	for {
		for _, bus := range buses {
			if timestamp%bus.id == (bus.id*100-bus.pos)%bus.id {
				buses = excludeBus(buses, bus)
				if len(buses) == 0 {
					return timestamp, nil
				}
				interval *= bus.id
			}
		}
		timestamp += interval
	}

	return 0, errors.New("Couldn't find matching timestamp")
}

func partOne(earliestDeparture int64, buses []Bus) int64 {
	var minBus Bus
	var minDiff int64

	minDiff = buses[0].id % earliestDeparture
	for _, bus := range buses {
		numDepartures := earliestDeparture / bus.id
		diff := bus.id * (numDepartures + 1) % earliestDeparture
		if diff < minDiff {
			minDiff = diff
			minBus = bus
		}
	}
	return minBus.id * minDiff
}

func main() {
	earliestDeparture, buses := getInput()
	earliestBusToAirport := partOne(earliestDeparture, buses)
	fmt.Println("The earliest bus that goes to the airport:", earliestBusToAirport)
	contestAnswer, _ := partTwo(buses)
	fmt.Println("The winning entry to the contest:", contestAnswer)
}
