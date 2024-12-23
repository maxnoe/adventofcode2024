package day22

import (
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


func Mix(n int, other int) int {
	return n ^ other
}

func Prune(n int) int {
	return n % 16777216
}

func Step(n int) int {
	n = Prune(Mix(n, n * 64))
	n = Prune(Mix(n, n / 32))
	n = Prune(Mix(n, n * 2048))
	return n
}

func Simulate(n int) int {
	for range 2000 {
		n = Step(n)
	}
	return n
}


func Parse(input string) ([]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numbers := make([]int, len(lines))

	var err error
	for i, line := range lines {
		numbers[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
	}

	return numbers, nil
}


func Part1(numbers []int) (int, error) {
	return aoc24.SumFunc(numbers, Simulate), nil
}

func Part2(numbers []int) (int, error) {
	return aoc24.SumFunc(numbers, Simulate), nil
}


func init() {
	aoc24.AddSolution(22, Parse, Part1, Part2)
}

