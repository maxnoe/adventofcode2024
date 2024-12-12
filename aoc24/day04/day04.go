package day04

import (
	"bufio"
	"slices"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

func Parse(input string) ([][]rune, error) {
	grid := make([][]rune, 0, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		grid = append(grid, []rune(strings.TrimSpace(scanner.Text())))
	}

	return grid, nil
}

var XMAS = []rune("XMAS")
var SAMX = []rune("SAMX")

const LEN = len("XMAS")

var START = XMAS[0]
var END = XMAS[LEN-1]

func isXMAS(substr []rune) bool {
	return slices.Equal(substr, XMAS) || slices.Equal(substr, SAMX)
}

func checkXMAS(grid [][]rune, i int, j int) int {
	n_rows := len(grid)
	n_cols := len(grid[0])
	n := 0

	c := grid[i][j]
	// if not start or end, we can already stop
	if c != 'X' && c != 'S' {
		return 0
	}

	// horizontal
	if j <= (n_cols - LEN) {
		if isXMAS(grid[i][j : j+LEN]) {
			n += 1
		}
	}

	substr := [LEN]rune{}
	// vertical
	if i <= (n_rows - LEN) {
		for k := range LEN {
			substr[k] = grid[i+k][j]
		}
		if isXMAS(substr[:]) {
			n += 1
		}
	}

	// diagonal 2
	if i <= (n_rows-LEN) && j <= (n_cols-LEN) {
		for k := range LEN {
			substr[k] = grid[i+k][j+k]
		}
		if isXMAS(substr[:]) {
			n += 1
		}
	}

	if i <= (n_rows-LEN) && j >= (LEN-1) {
		for k := range LEN {
			substr[k] = grid[i+k][j-k]
		}
		if isXMAS(substr[:]) {
			n += 1
		}
	}

	return n
}

func checkCross(grid [][]rune, i int, j int) bool {
	if grid[i+1][j+1] != 'A' {
		return false
	}

	// M . M
	// . A .
	// S . S
	if grid[i][j] == 'M' && grid[i+2][j] == 'M' && grid[i][j+2] == 'S' && grid[i+2][j+2] == 'S' {
		return true
	}

	// S . S
	// . A .
	// M . M
	if grid[i][j] == 'S' && grid[i+2][j] == 'S' && grid[i][j+2] == 'M' && grid[i+2][j+2] == 'M' {
		return true
	}

	// S . M
	// . A .
	// S . M
	if grid[i][j] == 'S' && grid[i+2][j] == 'M' && grid[i][j+2] == 'S' && grid[i+2][j+2] == 'M' {
		return true
	}

	// M . S
	// . A .
	// M . S
	if grid[i][j] == 'M' && grid[i+2][j] == 'S' && grid[i][j+2] == 'M' && grid[i+2][j+2] == 'S' {
		return true
	}

	return false
}

func Part1(grid [][]rune) (int, error) {
	n := 0
	n_rows := len(grid)
	n_cols := len(grid[0])
	for i := range n_rows {
		for j := range n_cols {
			n += checkXMAS(grid, i, j)
		}
	}
	return n, nil
}

func Part2(grid [][]rune) (int, error) {
	n := 0
	n_rows := len(grid)
	n_cols := len(grid[0])
	for i := range n_rows - 2 {
		for j := range n_cols - 2 {
			if checkCross(grid, i, j) {
				n += 1
			}
		}
	}
	return n, nil
}

func init() {
	aoc24.AddSolution(4, Parse, Part1, Part2)
}
