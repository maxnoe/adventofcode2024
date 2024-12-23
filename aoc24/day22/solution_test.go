package day22

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input = `
1
10
100
2024
`

func TestPrune(t *testing.T) {
	aoc24.AssertEqual(t, Prune(100000000), 16113920)
}

func TestMix(t *testing.T) {
	aoc24.AssertEqual(t, Mix(42, 15), 37)
}


func TestSimulate(t *testing.T) {
	aoc24.AssertEqual(t, Simulate(1), 8685429)
	aoc24.AssertEqual(t, Simulate(10), 4700978)
	aoc24.AssertEqual(t, Simulate(100), 15273692)
	aoc24.AssertEqual(t, Simulate(2024), 8667524)
}


func TestPart1(t *testing.T) {
	input, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	answer, err := Part1(input)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, answer, 37327623)
}
