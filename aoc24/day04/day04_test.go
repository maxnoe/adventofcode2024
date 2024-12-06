package day04

import (
	"strings"
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
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
	input, _ := ParseInput(day04_test_input)
	result, _ := Part1(input)
	aoc24.AssertEqual(t, result, 18)
}

func TestDay04Part2(t *testing.T) {
	input, _ := ParseInput(day04_test_input)
	result, _ := Part2(input)
	aoc24.AssertEqual(t, result, 9)
}
