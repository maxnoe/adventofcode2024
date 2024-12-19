package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Equation struct {
	Value   int
	Numbers []int
}

func Parse(input string) ([]Equation, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	equations := make([]Equation, 0, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Could not parse line %d: %s", i, line)
		}
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		numbers := strings.Fields(parts[1])
		equation := Equation{value, make([]int, len(numbers))}

		for i, number := range numbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			equation.Numbers[i] = value
		}

		equations = append(equations, equation)

	}
	return equations, nil
}

func IsValidRec(value int, current int, numbers []int, combine bool) bool {
	if len(numbers) == 0 {
		return value == current
	}

	if IsValidRec(value, current+numbers[0], numbers[1:], combine) {
		return true
	}
	if combine {
		if IsValidRec(value, Combine(current, numbers[0]), numbers[1:], combine) {
			return true
		}
	}
	return IsValidRec(value, current*numbers[0], numbers[1:], combine)
}

func IsValid(eq Equation, combine bool) bool {
	if IsValidRec(eq.Value, eq.Numbers[0]+eq.Numbers[1], eq.Numbers[2:], combine) {
		return true
	}
	if combine {
		if IsValidRec(eq.Value, Combine(eq.Numbers[0], eq.Numbers[1]), eq.Numbers[2:], combine) {
			return true
		}
	}
	return IsValidRec(eq.Value, eq.Numbers[0]*eq.Numbers[1], eq.Numbers[2:], combine)
}

func Combine(val1 int, val2 int) int {
	n := 1 + int(math.Log10(float64(val2)))
	return val1*int(math.Pow10(n)) + val2
}

func CountValid(eqs []Equation, combine bool) (int, error) {
	n := 0
	for _, eq := range eqs {
		if IsValid(eq, combine) {
			n += eq.Value
		}
	}

	return n, nil
}

func Part1(eqs []Equation) (int, error) {
	return CountValid(eqs, false)
}

func Part2(eqs []Equation) (int, error) {
	return CountValid(eqs, true)
}

func init() {
	aoc24.AddSolution(7, Parse, Part1, Part2)
}
