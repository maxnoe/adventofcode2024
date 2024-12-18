package day16

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input_1 = `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`

var test_input_2 = `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
`

func TestPart1(t *testing.T) {
	maze, err := Parse(test_input_1)
	aoc24.AssertEqual(t, err, nil)

	result, err := Part1(maze)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, result, 7036)
}

func TestPart2_1(t *testing.T) {
	maze, err := Parse(test_input_1)
	aoc24.AssertEqual(t, err, nil)

	result, err := Part2(maze)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, result, 45)
}

func TestPart2_2(t *testing.T) {
	maze, err := Parse(test_input_2)
	aoc24.AssertEqual(t, err, nil)

	result, err := Part2(maze)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, result, 64)
}
