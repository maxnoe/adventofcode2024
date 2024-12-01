package aoc24

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parseInput(input string) ([]int, []int, error) {
	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(strings.NewReader(input))
	n := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n += 1
		if line != "" {
			result := strings.Fields(line)
			if len(result) != 2 {
				msg := fmt.Sprintf("Expected 2 elements in line %d, got %d", n, len(result))
				return nil, nil, errors.New(msg)
			}

			valA, err := strconv.Atoi(result[0])
			if err != nil {
				return nil, nil, err
			}
			valB, err := strconv.Atoi(result[1])
			if err != nil {
				return nil, nil, err
			}

			list1 = append(list1, valA)
			list2 = append(list2, valB)
		}

	}

	return list1, list2, nil
}

func part1(list1 []int, list2 []int) (int, error) {

	result := 0
	for i, val1 := range list1 {
		val2 := list2[i]
		d := absDiff(val1, val2)
		result += d
	}

	return result, nil
}

func part2(list1 []int, list2 []int) (int, error) {

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
	for _, val1 := range list1 {

		if pos2 >= len(list2) {
			break
		}

		for list2[pos2] < val1 {
			pos2 += 1
			if pos2 >= len(list2) {
				break outer
			}
		}
		current_num_start = pos2
		count := 0
		for list2[pos2] == val1 {
			count += 1
			pos2 += 1
			if pos2 >= len(list2) {
				break
			}
		}
		pos2 = current_num_start

		result += val1 * count
	}

	return result, nil
}

func Day01(input string) error {
	list1, list2, err := parseInput(input)
	if err != nil {
		return err
	}

	// sort here, both parts profit
	slices.Sort(list1)
	slices.Sort(list2)

	start := time.Now()
	solution1, err := part1(list1, list2)
	stop := time.Now()
	if err != nil {
		return err
	}
	log.Printf("Part 1: %d in %d μs\n", solution1, stop.Sub(start).Microseconds())

	start = time.Now()
	solution2, err := part2(list1, list2)
	stop = time.Now()
	if err != nil {
		return err
	}
	log.Printf("Part 2: %d in %d μs\n", solution2, stop.Sub(start).Microseconds())

	return nil
}
