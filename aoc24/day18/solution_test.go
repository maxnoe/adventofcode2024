package day18

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var testInput = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0	
`

func TestPart1(t *testing.T) {
	positions, err := Parse(testInput)
	aoc24.AssertEqual(t, err, nil)

	aoc24.AssertEqual(t, findShortestPath(6, 6, positions[:12]), 22)
}

func TestPart2(t *testing.T) {
	positions, err := Parse(testInput)
	aoc24.AssertEqual(t, err, nil)

	pos, err := part2Impl(6, 6, positions)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, pos.X, 6)
	aoc24.AssertEqual(t, pos.Y, 1)
}
