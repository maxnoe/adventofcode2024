package day10

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestScore(t *testing.T) {
	grid, _ := Parse(test_input)

	aoc24.AssertEqual(t, Score(Pos{0, 2}, grid, false), 5)
	aoc24.AssertEqual(t, Score(Pos{0, 4}, grid, false), 6)
	aoc24.AssertEqual(t, Score(Pos{2, 4}, grid, false), 5)
}

func TestPart1(t *testing.T) {
	grid, _ := Parse(test_input)

	result, _ := Part1(grid)
	aoc24.AssertEqual(t, result, 36)
}

func TestRating(t *testing.T) {
	grid, _ := Parse(test_input)

	aoc24.AssertEqual(t, Score(Pos{0, 2}, grid, true), 20)
	aoc24.AssertEqual(t, Score(Pos{0, 4}, grid, true), 24)
	aoc24.AssertEqual(t, Score(Pos{2, 4}, grid, true), 10)
}

func TestPart2(t *testing.T) {
	grid, _ := Parse(test_input)

	result, _ := Part2(grid)
	aoc24.AssertEqual(t, result, 81)
}
