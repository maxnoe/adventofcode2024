package aoc24

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)


type Mul struct {
	left int
	right int
}


var mulRe = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

func FindMuls(input string) ([]Mul, error) {
	matches := mulRe.FindAllStringSubmatch(input, -1)
	muls := make([]Mul, len(matches))
	for i, match := range matches {
		left, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("Error parsing int: %v", match)
		}
		right, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("Error parsing int: %v", match)
		}
		muls[i] = Mul{left, right}
	}
	return muls, nil
}

func Day03Part1(input string) (int, error) {
	muls, err := FindMuls(input)
	if err != nil {
		return 0, err
	}
	result := SumFunc(muls, func(mul Mul) int {return mul.left * mul.right})
	return result, nil
}

func Day03(input string) error {
	start := time.Now()
	solution1, err := Day03Part1(input)
	stop := time.Now()
	if err != nil {
		return err
	}
	log.Printf("Part 1: %d in %d μs\n", solution1, stop.Sub(start).Microseconds())

	return nil
}

func init() {
	AddSolution(3, Day03)
}
