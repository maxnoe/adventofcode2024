package aoc24

import (
	"slices"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func AssertSliceEqual[T comparable](t *testing.T, actual []T, expected []T) {
	t.Helper()
	if !slices.Equal(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

type Integer interface {
	int | int8 | int16 | int32 | int64
}

func AbsDiff[I Integer](x, y I) I {
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

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	tmp := a
	a = b
	b = tmp % a
	return GCD(a, b)
}


func CopyAppend[T any](sl []T, pos T) []T {
	out := make([]T, len(sl)+1)
	copy(out, sl)
	out[len(out)-1] = pos
	return out
}

