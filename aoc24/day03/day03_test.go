package day03

import (
	"slices"
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var day03_test_input1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var day03_test_input2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

var expected_instructions1 = []Instruction{
	{MUL, []int{2, 4}},
	{MUL, []int{5, 5}},
	{MUL, []int{11, 8}},
	{MUL, []int{8, 5}},
}

var expected_instructions2 = []Instruction{
	{MUL, []int{2, 4}},
	{DONT, nil},
	{MUL, []int{5, 5}},
	{MUL, []int{11, 8}},
	{DO, nil},
	{MUL, []int{8, 5}},
}

func InstructionEqual(i1 Instruction, i2 Instruction) bool {
	if i1.typ != i2.typ {
		return false
	}

	return slices.Equal(i1.args, i2.args)
}

func TestFindInstructions(t *testing.T) {
	result, err := FindInstructions(day03_test_input1)

	aoc24.AssertEqual(t, err, nil)

	if !slices.EqualFunc(result, expected_instructions1, InstructionEqual) {
		t.Fatalf("unexpected muls, got %v, expeceted %v", result, expected_instructions1)
	}

	result, err = FindInstructions(day03_test_input2)

	aoc24.AssertEqual(t, err, nil)

	if !slices.EqualFunc(result, expected_instructions2, InstructionEqual) {
		t.Fatalf("unexpected muls, got %v, expeceted %v", result, expected_instructions2)
	}
}

func TestPart01(t *testing.T) {
	result, _ := Part1(expected_instructions1)
	aoc24.AssertEqual(t, result, 161)
}

func TestPart02(t *testing.T) {
	result, _ := Part2(expected_instructions2)
	aoc24.AssertEqual(t, result, 48)
}
