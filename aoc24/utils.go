package aoc24


func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
