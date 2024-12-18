package day16

import (
	"log"
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

var Vecs = []Pos {
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type State struct {
	P Pos
	Dir Direction
}

type PathHead struct {
	S State
	Cost int
	History []Pos
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

func CopyAppend(path []Pos, pos Pos) []Pos {
	out := make([]Pos, len(path) + 1)
	copy(out, path)
	out[len(out) - 1] = pos
	return out
}

func FindBestPaths(maze Maze) []PathHead {
	visited := make(map[State]int)

	state := State{maze.Start, EAST}
	start := PathHead{state, 0, []Pos{maze.Start}}
	to_check := []PathHead{start}

	paths := make([]PathHead, 0)
	min_cost := math.MaxInt
	for len(to_check) > 0 {
		head := to_check[0]
		to_check = to_check[1:]
		pos := head.S.P

		if pos.R == maze.End.R && pos.C == maze.End.C {
			if head.Cost <= min_cost {
				min_cost = head.Cost
				paths = append(paths, head)
			}
			continue
		}

		cost, found := visited[head.S]
		if !found || head.Cost <= cost {
			visited[head.S] = head.Cost
		}
		if found && head.Cost > cost {
			continue
		}

		for dir, vec := range Vecs {
			dir := Direction(dir)
			n := Pos{pos.R + vec.R, pos.C + vec.C}
			
			// out of bounds check
			if n.R < 0 || n.R >= maze.Rows || n.C < 0 || n.C >= maze.Cols {
				continue
			}
			// check for wall
			if maze.Grid[n.R][n.C] {
				continue
			}

			cost := Cost(head.S.Dir, dir)
			proposal := PathHead{State{n, dir}, head.Cost + cost, CopyAppend(head.History, n)}
			to_check = append(to_check, proposal)
		}
	}

	best_paths := make([]PathHead, 0)
	for _, path := range paths {
		if path.Cost == min_cost {
			best_paths = append(best_paths, path)
		}
	}
	return best_paths
}

func PrintPaths(maze Maze, paths []PathHead) {
	lines := make([][]rune, maze.Rows)
	for row := range maze.Rows {
		line := make([]rune, maze.Cols)
		for col := range maze.Cols {
			if maze.Grid[row][col] {
				line[col] = '#'
			} else if maze.Start.R == row && maze.Start.C == col {
				line[col] = 'S'
			} else if maze.End.R == row && maze.End.C == col {
				line[col] = 'E'
			} else {
				line[col] = '.'
			}
		}
		lines[row] = line
	}

	for _, path := range paths {
		for _, p := range path.History {
			if lines[p.R][p.C] == '.' {
				lines[p.R][p.C] = 'O'
			}
		}
	}

	for _, line := range lines {
		log.Println(string(line))
	}
}

func Part1(maze Maze) (int, error) {
	paths := FindBestPaths(maze)
	return paths[0].Cost, nil
}

func Part2(maze Maze) (int, error) {
	spots := make(map[Pos]struct{})
	paths := FindBestPaths(maze)
	for _, path := range paths {
		for _, pos := range path.History {
			spots[pos] = struct{}{}
		}
	}
	return len(spots), nil
}

func init() {
	aoc24.AddSolution(16, Parse, Part1, Part2)
}
