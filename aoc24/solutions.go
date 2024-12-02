package aoc24

import "log"

var solutions = make(map[int]func(string) error)

func AddSolution(day int, f func(string) error) {
	solutions[day] = f
}

func Solve(day int, input string) error {
	f := solutions[day]
	if f == nil {
		log.Fatalf("Day %d is not yet implemented", day)
	}

	return f(input)
}
