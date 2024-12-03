package aoc24

import (
	"slices"
	"testing"
)


var day03_test_input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

var expected_muls = []Mul {
	{2, 4},
	{5, 5},
	{11, 8},
	{8, 5},
}

func TestFindMuls(t *testing.T) {
	result, err := FindMuls(day03_test_input)

	assertEqual(t, err, nil)
	assertEqual(t, len(result), 4)

	if !slices.Equal(result, expected_muls) {
		t.Fatalf("unexpected muls, got %v, expeceted %v", result, expected_muls)
	}
}

func TestPart01(t *testing.T) {
	result, err := Day03Part1(day03_test_input)
	assertEqual(t, err, nil)
	assertEqual(t, result, 161)
}
