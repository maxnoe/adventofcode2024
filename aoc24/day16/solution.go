package day16

import (
	"math"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Pos struct {
	R int
	C int
}

type Maze struct {
	Grid  [][]bool
	Start Pos
	End   Pos
	Rows  int
	Cols  int
}

func Parse(input string) (Maze, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	maze := Maze{}
	for i, line := range lines {
		maze.Grid = append(maze.Grid, make([]bool, len(line)))
		for j, chr := range line {
			if chr == '#' {
				maze.Grid[i][j] = true
			} else if chr == 'S' {
				maze.Start = Pos{i, j}
			} else if chr == 'E' {
				maze.End = Pos{i, j}
			}
		}
	}
	maze.Rows = len(maze.Grid)
	maze.Cols = len(maze.Grid[0])
	return maze, nil
}

type Direction int8

const (
	EAST  Direction = 0
	SOUTH           = 1
	WEST            = 2
	NORTH           = 3
)

var Vecs = map[Direction]Pos {
	EAST: {0, 1},
	SOUTH: {1, 0},
	WEST: {0, -1},
	NORTH: {-1, 0},
}

type PathHead struct {
	P    Pos
	Cost int
	Dir  Direction
}

func Cost(current Direction, wanted Direction) int {
	if current == wanted {
		return 1
	}
	diff := aoc24.AbsDiff(int8(current), int8(wanted))
	if diff == 1 {
		return 1001
	}
	if diff == 2 {
		return 2001
	}
	// diff = 3 means EAST/North => 1 turn
	return 1001
}

func Part1(maze Maze) (int, error) {
	visited := make([][]int, maze.Rows)
	for i := range maze.Rows {
		visited[i] = make([]int, maze.Cols)
		for j := range maze.Cols {
			visited[i][j] = math.MaxInt
		}
	}

	to_check := []PathHead{{maze.Start, 0, EAST}}

	for len(to_check) > 0 {
		head := to_check[0]
		to_check = to_check[1:]

		for dir, vec := range Vecs {
			n:= Pos{head.P.R + vec.R, head.P.C + vec.C}
			
			// out of bounds check
			if n.R < 0 || n.R >= maze.Rows || n.C < 0 || n.C >= maze.Cols {
				continue
			}
			// check for wall
			if maze.Grid[n.R][n.C] {
				continue
			}

			cost := Cost(head.Dir, dir)
			proposal := PathHead{n, head.Cost + cost, dir}

			if visited[n.R][n.C] > proposal.Cost {
				visited[n.R][n.C] = proposal.Cost

				if n.R != maze.End.R || n.C != maze.End.C {
					to_check = append(to_check, proposal)
				} 
			}
		}
	}
	return visited[maze.End.R][maze.End.C], nil
}

func Part2(maze Maze) (int, error) {
	return 0, nil
}

func init() {
	aoc24.AddSolution(16, Parse, Part1, Part2)
}
