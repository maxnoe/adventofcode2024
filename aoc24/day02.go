package aoc24

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Day02ParseInput(input string) ([][]int, error) {
	data := make([][]int, 0, 50)
	scanner := bufio.NewScanner(strings.NewReader(input))
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tokens := strings.Fields(line)
		data = append(data, make([]int, len(tokens)))
		for j, token := range tokens {
			val, err := strconv.Atoi(token)
			if err != nil {
				return nil, fmt.Errorf("Error parsing int in line %d, column %d", i, j)
			}
			data[i][j] = val
		}

		i += 1
	}

	return data, nil
}

func Day02ReportSafe(report []int) bool {
	ascending := true
	previous := report[0]
	for i, val := range report[1:] {
		if diff := absDiff(val, previous); diff == 0 || diff > 3 {
			return false
		}

		if i == 0 {
			ascending = val > previous
		} else {
			if ascending && val < previous {
				return false
			}
			if !ascending && val > previous {
				return false
			}
		}
		previous = val
	}
	return true
}

func Day02Part1(reports [][]int) int {
	return CountTrueFunc(reports, Day02ReportSafe)
}

func Day02ReportSafeWithDampener(report []int) bool {
	if Day02ReportSafe(report) {
		return true
	}

	for i := range len(report) {
		dampened_report := make([]int, 0, len(report)-1)
		dampened_report = append(dampened_report, report[:i]...)
		dampened_report = append(dampened_report, report[i+1:]...)
		if Day02ReportSafe(dampened_report) {
			return true
		}
	}

	return false
}

func Day02Part2(reports [][]int) int {
	return CountTrueFunc(reports, Day02ReportSafeWithDampener)
}

func Day02(input string) error {
	reports, err := Day02ParseInput(input)
	if err != nil {
		return err
	}

	start := time.Now()
	solution1 := Day02Part1(reports)
	stop := time.Now()
	log.Printf("Part 1: %d in %d μs\n", solution1, stop.Sub(start).Microseconds())

	start = time.Now()
	solution2 := Day02Part2(reports)
	stop = time.Now()
	log.Printf("Part 2: %d in %d μs\n", solution2, stop.Sub(start).Microseconds())

	return nil
}

func init() {
	AddSolution(2, Day02)
}
