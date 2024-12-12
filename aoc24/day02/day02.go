package day02

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

func Parse(input string) ([][]int, error) {
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

func ReportSafe(report []int) bool {
	ascending := true
	previous := report[0]
	for i, val := range report[1:] {
		if diff := aoc24.AbsDiff(val, previous); diff == 0 || diff > 3 {
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

func Part1(reports [][]int) (int, error) {
	return aoc24.CountTrueFunc(reports, ReportSafe), nil
}

func ReportSafeWithDampener(report []int) bool {
	if ReportSafe(report) {
		return true
	}

	for i := range len(report) {
		dampened_report := make([]int, 0, len(report)-1)
		dampened_report = append(dampened_report, report[:i]...)
		dampened_report = append(dampened_report, report[i+1:]...)
		if ReportSafe(dampened_report) {
			return true
		}
	}

	return false
}

func Part2(reports [][]int) (int, error) {
	return aoc24.CountTrueFunc(reports, ReportSafeWithDampener), nil
}

func init() {
	aoc24.AddSolution(2, Parse, Part1, Part2)
}
