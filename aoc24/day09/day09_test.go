package day09

import (
	"slices"
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = "2333133121414131402"
var disk_space = []int8{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}

func TestParse(t *testing.T) {
	result, err := Parse(test_input)

	aoc24.AssertEqual(t, err, nil)
	if !slices.Equal(result, disk_space) {
		t.Fatalf("Did not match: %v != %v", result, disk_space)
	}
}

func TestPart1(t *testing.T) {
	res, err := Part1(disk_space)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, res, 1928)

}
