package aoc24

import (
	"strings"
	"testing"
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
	lab, guard := Day06ParseInput(test_input_06)

	assertEqual(t, Day06Part1(lab, guard), 41)
}
