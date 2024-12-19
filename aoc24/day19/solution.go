package day19

import (
	"fmt"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type OnsenWishlist struct {
	Patterns []string
	Towels   []string
}

func Parse(input string) ([]int, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid format of input, expected two parts separated by newline")
	}
	patterns := strings.Split(parts[0], ", ")
	towels := strings.Split(parts[1], "\n")
	possibilites := make([]int, len(towels))
	for i, t := range towels {
		possibilites[i] = CountPossibilities(t, patterns)
	}
	return possibilites, nil
}

func CountPossibilitiesMatches(towel string, matches map[int][]string, index int, cache map[int]int) int {
	if val, found := cache[index]; found {
		return val
	}

	patterns, found := matches[index]
	if !found {
		return 0
	}

	total := 0
	for _, pattern := range patterns {
		next_index := index + len(pattern)

		// we found a solution
		if next_index == len(towel) {
			total += 1
			continue
		}

		n := CountPossibilitiesMatches(towel, matches, next_index, cache)
		total += n
	}
	cache[index] = total
	return total
}

func CountPossibilities(towel string, patterns []string) int {
	matches := make(map[int][]string)

	for _, pattern := range patterns {

		start := 0
		for start < len(towel) {
			offset := strings.Index(towel[start:], pattern)
			if offset == -1 {
				break
			}

			index := start + offset
			matches[index] = append(matches[index], pattern)
			start = index + 1
		}
	}

	return CountPossibilitiesMatches(towel, matches, 0, make(map[int]int))
}

func Part1(possibilities []int) (int, error) {
	n := aoc24.CountTrueFunc(possibilities, func(n int) bool { return n != 0})
	return n, nil
}

func Part2(possibilities []int) (int, error) {
	n := aoc24.Sum(possibilities)
	return n, nil
}

func init() {
	aoc24.AddSolution(19, Parse, Part1, Part2)
}
