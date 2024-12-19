package day19

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func TestCount(t *testing.T) {
	wishlist, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	expected := []int{2, 1, 4, 6, 0, 1, 2, 0}

	for i, n := range expected {
		if r := CountPossibilities(wishlist.Towels[i], wishlist.Patterns); r != n {
			t.Errorf("Expected towel %s to be %d, got %d", wishlist.Towels[i], n, r)
		}
	}
}

func TestPart1(t *testing.T) {
	wishlist, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	answer, err := Part1(wishlist)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, answer, 6)
}

func TestPart2(t *testing.T) {
	wishlist, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	answer, err := Part2(wishlist)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, answer, 16)
}
