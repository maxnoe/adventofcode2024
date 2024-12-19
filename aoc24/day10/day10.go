package day10

import (
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

func Parse(input string) ([][]int8, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	topography := make([][]int8, len(lines))

	for i, line := range lines {
		topography[i] = make([]int8, len(line))
		for j, chr := range line {
			topography[i][j] = int8(chr - '0')
		}
	}

	return topography, nil
}

type Pos struct {
	R int
	C int
}

var DIRECTIONS = [4]Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func Score(pos Pos, topography [][]int8, trails bool) int {
	n := 0

	visited := make(map[Pos]struct{})

	to_check := make([]Pos, 0)
	to_check = append(to_check, pos)

	n_rows := len(topography)
	n_cols := len(topography[1])

	for len(to_check) > 0 {
		pos = to_check[0]
		to_check = to_check[1:]
		if _, v := visited[pos]; v && !trails {
			continue
		}

		visited[pos] = struct{}{}

		level := topography[pos.R][pos.C]
		if level == 9 {
			n += 1
			continue
		}

		for _, d := range DIRECTIONS {
			row := pos.R + d.R
			col := pos.C + d.C

			// boundary checks
			if row < 0 || row >= n_rows || col < 0 || col >= n_cols {
				continue
			}

			new_pos := Pos{row, col}
			if topography[row][col] == (level + 1) {
				to_check = append(to_check, new_pos)
			}
		}
	}

	return n
}

func TotalScore(topography [][]int8, trails bool) int {
	result := 0

	for i, row := range topography {
		for j, val := range row {
			if val == 0 {
				result += Score(Pos{i, j}, topography, trails)
			}
		}
	}
	return result
}

func Part1(topography [][]int8) (int, error) {
	return TotalScore(topography, false), nil
}

func Part2(topography [][]int8) (int, error) {
	return TotalScore(topography, true), nil
}

func init() {
	aoc24.AddSolution(10, Parse, Part1, Part2)
}
