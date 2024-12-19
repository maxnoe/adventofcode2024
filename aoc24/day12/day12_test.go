package day12

import (
	"reflect"
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`

var expected = [][]int{
	{1, 1, 1, 1, 2, 2, 3, 3, 4, 4},
	{1, 1, 1, 1, 2, 2, 3, 3, 3, 4},
	{5, 5, 1, 1, 1, 3, 3, 4, 4, 4},
	{5, 5, 1, 3, 3, 3, 6, 4, 4, 4},
	{5, 5, 5, 5, 3, 6, 6, 7, 4, 8},
	{5, 5, 9, 5, 3, 3, 6, 6, 8, 8},
	{5, 5, 9, 9, 9, 3, 6, 6, 8, 8},
	{10, 9, 9, 9, 9, 9, 6, 6, 8, 8},
	{10, 9, 9, 9, 11, 9, 6, 8, 8, 8},
	{10, 10, 10, 9, 11, 11, 6, 8, 8, 8},
}

func TestFindRegions(t *testing.T) {
	grid, _ := Parse(test_input)
	regions, region_map := FindRegions(grid)
	aoc24.AssertEqual(t, len(regions), 11)
	if !reflect.DeepEqual(expected, region_map) {
		t.Fatalf("%v", region_map)
	}

	aoc24.AssertEqual(t, regions[0].Perimeter, 18)
	aoc24.AssertEqual(t, regions[0].Area, 12)
	aoc24.AssertEqual(t, regions[0].Plant, 'R')
}

func TestPart1(t *testing.T) {
	input, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	result, err := Part1(input)
	aoc24.AssertEqual(t, result, 1930)
}
