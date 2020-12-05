package main

import (
	"testing"
)

func TestGetSeatId(t *testing.T) {
	var tests = map[string]int{
		"FBFBBFFRLR": 357,
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
	}

	for boardingPass, expectedSeatId := range tests {
		actualSeatId := getSeatId(boardingPass)
		if expectedSeatId != actualSeatId {
			t.Logf("Expected %d, got %d", expectedSeatId, actualSeatId)
			t.Fail()
		}
	}
}
