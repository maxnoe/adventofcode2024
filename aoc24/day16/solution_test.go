package day16

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input = `
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

func TestPart1(t *testing.T) {
	maze, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	result, err := Part1(maze)
	aoc24.AssertEqual(t, result, 7036)
}
