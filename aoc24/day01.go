package aoc24

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int, error) {
	var listOne []int
	var listTwo []int
	
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
			if err != nil {return nil, nil, err}
			valB, err := strconv.Atoi(result[1])
			if err != nil {return nil, nil, err}

			listOne = append(listOne, valA)
			listTwo = append(listTwo, valB)
		}
		
	}

	return listOne, listTwo, nil
}

func part1(listOne []int, listTwo []int) (int, error) {
	sort.Ints(listOne)
	sort.Ints(listTwo)

	result := 0
	for i, val1 := range listOne {
		val2 := listTwo[i]
		d := absDiff(val1, val2)
		result += d
		log.Printf("%d %d %d %d\n", val1, val2, d, result)
	}

	return result, nil
}

func Day01(input string) error {
	listOne, listTwo, err := parseInput(input)
	if err != nil {return err}

	val, err := part1(listOne, listTwo)
	if err != nil {return err}
	log.Printf("Part 1: %d\n", val);

	return nil
}
