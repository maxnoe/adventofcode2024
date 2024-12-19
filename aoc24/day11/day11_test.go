package day11

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `125 17`

func TestEvolve(t *testing.T) {
	stones, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	aoc24.AssertEqual(t, NumStones(stones, 0), 2)
	aoc24.AssertEqual(t, NumStones(stones, 1), 3)
	aoc24.AssertEqual(t, NumStones(stones, 2), 4)
	aoc24.AssertEqual(t, NumStones(stones, 3), 5)
	aoc24.AssertEqual(t, NumStones(stones, 4), 9)
	aoc24.AssertEqual(t, NumStones(stones, 5), 13)
	aoc24.AssertEqual(t, NumStones(stones, 6), 22)
	aoc24.AssertEqual(t, NumStones(stones, 25), 55312)
}
