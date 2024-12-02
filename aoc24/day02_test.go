package aoc24

import (
	"fmt"
	"slices"
	"testing"
)


var test_input_string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

var test_input = [][]int {
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}


func TestDay02ParseInput(t *testing.T) {
	input, err := Day02ParseInput(test_input_string)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(input) != len(test_input) {
		t.Fatalf("Wrong number of elements, got %d, expected %d", len(input), len(test_input))
	}

	for i, line := range test_input {
		if !slices.Equal(line, input[i]) {
			t.Fatalf("Parsed input does not match expectation in line %d: %v != %v", i, line, input[i])
		}
	}
}

func TestDay02Part1(t *testing.T) {
	expected := []bool{true, false, false, false, false, true}

	for i, report := range test_input {
		if Day02ReportSafe(report) != expected[i] {
			t.Log(fmt.Sprintf("Expected report %d to be %t", i, expected[i]))
			t.Fail()
		}
	}

	if answer := Day02Part1(test_input); answer != 2 {
		t.Fatalf("Expected 2 safe reports, got %d", answer)
	}

}
