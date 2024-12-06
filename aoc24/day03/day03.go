package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type InstructionType int8

const (
	MUL InstructionType = iota
	DO
	DONT
)

type Instruction struct {
	typ  InstructionType
	args []int
}

const doStr = "do()"
const dontStr = "don't()"
const lenDo = len(doStr)
const lenDont = len(dontStr)

var mulRe = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

func FindInstructions(input string) ([]Instruction, error) {
	instructions := make([]Instruction, 0)
	i := 0
	for i < len(input) {
		last := min(i+lenDo, len(input))
		if input[i:last] == doStr {
			instructions = append(instructions, Instruction{DO, nil})
			i += lenDo
			continue
		}

		last = min(i+lenDont, len(input))
		if input[i:last] == dontStr {
			instructions = append(instructions, Instruction{DONT, nil})
			i += lenDont
			continue
		}

		last = min(i+4, len(input))
		if input[i:last] == "mul(" {
			last := min(i+12, len(input))
			match := mulRe.FindStringSubmatch(input[i:last])
			if match != nil {
				left, err := strconv.Atoi(match[1])
				if err != nil {
					return nil, fmt.Errorf("Error parsing int: %v", match)
				}
				right, err := strconv.Atoi(match[2])
				if err != nil {
					return nil, fmt.Errorf("Error parsing int: %v", match)
				}
				instructions = append(instructions, Instruction{MUL, []int{left, right}})
				i += len(match[0])
			} else {
				i += 4
			}
			continue
		}

		i += 1
	}

	return instructions, nil
}

func instructionValue(i Instruction) int {
	if i.typ != MUL {
		return 0
	}
	return i.args[0] * i.args[1]
}

func Part1(instructions []Instruction) (int, error) {
	return aoc24.SumFunc(instructions, instructionValue), nil
}

func Part2(instructions []Instruction) (int, error) {
	result := 0
	do := 1
	for _, instr := range instructions {
		switch instr.typ {
		case DO:
			do = 1
		case DONT:
			do = 0
		case MUL:
			result += do * instr.args[0] * instr.args[1]
		}
	}
	return result, nil
}

func Day03(input string) error {
	return aoc24.Solve(input, FindInstructions, Part1, Part2)
}

func init() {
	aoc24.AddSolution(3, Day03)
}
