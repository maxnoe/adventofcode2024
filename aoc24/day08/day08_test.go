package day08

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestGCD(t *testing.T) {
	aoc24.AssertEqual(t, GCD(7, 14), 7)
	aoc24.AssertEqual(t, GCD(21, 14), 7)
}

func TestPart1(t *testing.T) {
	world, _ := Parse(test_input)
	res, _ := Part1(world)
	aoc24.AssertEqual(t, res, 14)
}

func TestPart2(t *testing.T) {
	world, _ := Parse(test_input)
	res, _ := Part2(world)
	aoc24.AssertEqual(t, res, 34)
}
