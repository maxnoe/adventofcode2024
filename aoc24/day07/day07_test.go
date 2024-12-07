package day07

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`


func TestPart1(t *testing.T) {
	eqs, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)
	result, err := Part1(eqs)
	aoc24.AssertEqual(t, result, 3749)
}


func TestCombiner(t *testing.T) {
	aoc24.AssertEqual(t, Combine(15, 6), 156)
	aoc24.AssertEqual(t, Combine(1234, 56789), 123456789)
	aoc24.AssertEqual(t, Combine(100, 10), 10010)
	aoc24.AssertEqual(t, Combine(10, 100), 10100)
}

func TestPart2(t *testing.T) {
	eqs, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)
	result, err := Part2(eqs)
	aoc24.AssertEqual(t, result, 11387)
}

