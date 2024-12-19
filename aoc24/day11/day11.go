package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type pair struct {
	a int
	b int
}

var cache = make(map[pair]int)

func Parse(input string) ([]int, error) {
	fields := strings.Fields(strings.TrimSpace(input))
	stones := make([]int, len(fields))
	for i, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		stones[i] = val
	}
	return stones, nil
}

func NumDigits(i int) int {
	return int(math.Log10(float64(i)) + 1)
}

func Split(stone int, digits int) pair {
	p := int(math.Pow10(digits / 2))
	return pair{stone / p, stone % p}
}

func Evolve(stone int, blinks int) int {
	key := pair{stone, blinks}
	if val, ok := cache[key]; ok {
		return val
	}

	if blinks == 0 {
		return 1
	}

	result := 0
	if stone == 0 {
		result = Evolve(1, blinks-1)
	} else if digits := NumDigits(stone); digits%2 == 0 {
		split := Split(stone, digits)
		result = Evolve(split.a, blinks-1) + Evolve(split.b, blinks-1)
	} else {
		result = Evolve(2024*stone, blinks-1)
	}

	cache[key] = result
	return result

}

func NumStones(stones []int, blinks int) int {
	total := 0
	for _, stone := range stones {
		total += Evolve(stone, blinks)
	}
	return total
}

func Part1(stones []int) (int, error) {
	return NumStones(stones, 25), nil
}

func Part2(stones []int) (int, error) {
	return NumStones(stones, 75), nil
}

func init() {
	aoc24.AddSolution(11, Parse, Part1, Part2)
}
