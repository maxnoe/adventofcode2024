package day17

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

func TestInstructions1(t *testing.T) {
	m := Machine{0, 0, 9, []int{2,6}, 0, nil}
	m.RunProgram()
	aoc24.AssertEqual(t, m.B, 1)
}

func TestInstructions2(t *testing.T) {
	m := Machine{10, 0, 0, []int{5,0,5,1,5,4}, 0, nil}
	m.RunProgram()
	aoc24.AssertSliceEqual(t, m.Output, []int{0,1,2})
}

func TestInstructions3(t *testing.T) {
	m := Machine{0, 29, 0, []int{1,7}, 0, nil}
	m.RunProgram()
	aoc24.AssertEqual(t, m.B, 26)
}

func TestInstructions4(t *testing.T) {
	m := Machine{2024, 0, 0, []int{0,1,5,4,3,0}, 0, nil}
	m.RunProgram()
	aoc24.AssertSliceEqual(t, m.Output, []int{4,2,5,6,7,7,7,7,3,1,0})
	aoc24.AssertEqual(t, m.A, 0)
}

func TestPart1(t *testing.T) {
	m, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, m.A, 729)
	aoc24.AssertEqual(t, m.B, 0)
	aoc24.AssertEqual(t, m.C, 0)
	aoc24.AssertSliceEqual(t, m.Program, []int{0,1,5,4,3,0})


	out, err := Part1(m)
	aoc24.AssertEqual(t, out, "4,6,3,5,6,3,5,2,1,0")
}


func TestPart2(t *testing.T) {
	m := Machine{2024, 0, 0, []int{0,3,5,4,3,0}, 0, nil}
	out, _ := Part2(m)
	aoc24.AssertEqual(t, out, 117440)
}
