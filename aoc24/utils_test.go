package aoc24

import "testing"

func assertEqual[T comparable] (t *testing.T, actual T, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestAbsDiff(t *testing.T)  {
	assertEqual(t, absDiff(0, 0), 0);
	assertEqual(t, absDiff(3, 1), 2);
	assertEqual(t, absDiff(-1, 5), 6);
}

func TestSumFuncInt(t *testing.T) {
	ints := []int{1, 2, 3}

	r := SumFunc(ints, func(val int) int {return val})
	assertEqual(t, r, 6);
}

func TestSumFuncFloat(t *testing.T) {
	floats := []float64{1.0, 2.0, 3.0}
	r := SumFunc(floats, func(val float64) float64 {return val * val})
	assertEqual(t, r, 14.0);
}

func TestSumFuncMix(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	r := SumFunc(ints, func(val int) float64 {return 0.5 * float64(val)})
	assertEqual(t, r, 7.5);
}

func TestCountTrueFuncInt(t *testing.T) {
	vals := []int{1, 2, 3}

	r := CountTrueFunc(vals, func(val int) bool {return val >= 2})
	assertEqual(t, r, 2);
}

func TestCountTrueFuncFloat(t *testing.T) {
	floats := []float64{1.0, 2.0, 2.5, 2.6, 3.0}

	r := CountTrueFunc(floats, func(val float64) bool {return val > 2.4})
	assertEqual(t, r, 3);
}
