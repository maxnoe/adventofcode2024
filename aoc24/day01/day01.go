package day01

import (
	"bufio"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Input struct {
	list1 []int
	list2 []int
}

func Parse(input string) (Input, error) {
	result := Input{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	n := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n += 1
		if line != "" {
			fields := strings.Fields(line)
			if len(fields) != 2 {
				msg := fmt.Sprintf("Expected 2 elements in line %d, got %d", n, len(fields))
				return Input{}, errors.New(msg)
			}

			valA, err := strconv.Atoi(fields[0])
			if err != nil {
				return Input{}, err
			}
			valB, err := strconv.Atoi(fields[1])
			if err != nil {
				return Input{}, err
			}

			result.list1 = append(result.list1, valA)
			result.list2 = append(result.list2, valB)
		}
	}
	slices.Sort(result.list1)
	slices.Sort(result.list2)
	return result, nil
}

func Part1(input Input) (int, error) {

	result := 0
	for i, val1 := range input.list1 {
		val2 := input.list2[i]
		d := aoc24.AbsDiff(val1, val2)
		result += d
	}

	return result, nil
}

func Part2(input Input) (int, error) {

	result := 0

	// much simpler, but also much slower
	// counts := make(map[int]int)
	// for _, val2 := range list2 {
	// 	counts[val2] += 1
	// }
	// for _, val1 := range list1 {
	// 	result += val1 * counts[val1]
	// }

	// doesn't allocate and relies on the list being sorted
	pos2 := 0
	current_num_start := 0

outer:
	for _, val1 := range input.list1 {

		if pos2 >= len(input.list2) {
			break
		}

		for input.list2[pos2] < val1 {
			pos2 += 1
			if pos2 >= len(input.list2) {
				break outer
			}
		}
		current_num_start = pos2
		count := 0
		for input.list2[pos2] == val1 {
			count += 1
			pos2 += 1
			if pos2 >= len(input.list2) {
				break
			}
		}
		pos2 = current_num_start

		result += val1 * count
	}

	return result, nil
}

func init() {
	aoc24.AddSolution(1, Parse, Part1, Part2)
}
