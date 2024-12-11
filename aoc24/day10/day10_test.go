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

	aoc24.AssertEqual(t, Score(Pos{0, 2}, grid), 5)
	aoc24.AssertEqual(t, Score(Pos{0, 4}, grid), 6)
	aoc24.AssertEqual(t, Score(Pos{2, 4}, grid), 5)
}

func TestPart1(t *testing.T) {
	grid, _ := Parse(test_input)

	result, _ := Part1(grid)
	aoc24.AssertEqual(t, result, 36)
}
