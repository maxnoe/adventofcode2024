package aoc24

import (
	"fmt"
	"log"
	"time"
)

var solutions = make(map[int]func(string) error)


func SolveDay(day int, input string) error {
	f := solutions[day]
	if f == nil {
		log.Fatalf("Day %d is not yet implemented", day)
	}

	return f(input)
}

func Solve[I any, O1 any, O2 any](input_string string, parse func(string) (I, error), part1 func(I) (O1, error), part2 func(I) (O2, error)) error {
	input, err := parse(input_string)
	if err != nil {
		return fmt.Errorf("Error parsing input: %v", err)
	}

	start := time.Now()
	solution1, err := part1(input)
	stop := time.Now()
	if err != nil {
		return fmt.Errorf("Error solving part 1: %v", err)
	}
	log.Printf("Part 1: %v in %d μs\n", solution1, stop.Sub(start).Microseconds())

	start = time.Now()
	solution2, err := part2(input)
	stop = time.Now()
	if err != nil {
		return fmt.Errorf("Error solving part 2: %v", err)
	}
	log.Printf("Part 2: %v in %d μs\n", solution2, stop.Sub(start).Microseconds())

	return nil
}

func AddSolution[I any, O1 any, O2 any](day int, parse func(string) (I, error), part1 func(I) (O1, error), part2 func(I) (O2, error)) {
	solutions[day] = func (input string) error {
		return Solve(input, parse, part1, part2)
	}
}
