package day06

import (
	"strings"
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input_06 = strings.TrimSpace(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`)


func TestDay06Part1(t *testing.T) {
	input, _ := Parse(test_input_06)

	result, _ := Part1(input)
	aoc24.AssertEqual(t, result , 41)
}

func TestDay06Part2(t *testing.T) {
	input, _ := Parse(test_input_06)

	result, _ := Part2(input)
	aoc24.AssertEqual(t, result, 6)
}
