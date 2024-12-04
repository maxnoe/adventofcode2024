package aoc24

import (
	"strings"
	"testing"
)

var day04_test_input = strings.TrimSpace(`
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`)

func TestDay04Part1(t *testing.T) {
	input := Day04ParseInput(day04_test_input)
	assertEqual(t, Day04Part1(input), 18)
}

func TestDay04Part2(t *testing.T) {
	input := Day04ParseInput(day04_test_input)
	assertEqual(t, Day04Part2(input), 9)
}
