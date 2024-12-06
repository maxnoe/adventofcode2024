package aoc24

import (
	"testing"
)

func AssertEqual[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}


func AbsDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

type Number interface {
	int | float32 | float64
}

func SumFunc[T any, V Number](inputs []T, f func(T) V) V {
	var sum V
	for _, input := range inputs {
		sum += f(input)
	}
	return sum
}

func CountTrueFunc[T any](inputs []T, f func(T) bool) int {
	count := 0
	for _, input := range inputs {
		if f(input) {
			count += 1
		}
	}
	return count
}
