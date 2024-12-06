package aoc24

import (
	"log"
	"strings"
	"time"
)


type GuardDirection int8

const (
	UP GuardDirection = '^'
	DOWN = 'v'
	LEFT = '<'
	RIGHT = '>'
)

type Lab [][]bool 

type Guard struct {
	row int
	col int
	direction GuardDirection
}

func Day06ParseInput(input string) (Lab, Guard) {
	guard := Guard{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	lab := make(Lab, len(lines))

	for i, line := range lines {
		lab[i] = make([]bool, len(line))
		for j, chr :=  range line {
			switch chr {
			case '#':
				lab[i][j] = true
			case '^':
				guard = Guard{i, j, UP}
			}
		}
	}

	return lab, guard
}


func countTrue(lab Lab) int {
	n := 0
	for _, row := range lab {
		for _, val := range row {
			if val {
				n += 1
			}
		}
	}
	return n
}


func nextDirection(dir GuardDirection) GuardDirection {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN: 
		return LEFT
	case LEFT:
		return UP
	}
	panic("Invalid direction")
}


func visitedFields(obstructions Lab, guard Guard) Lab {
	n_rows := len(obstructions)
	n_cols := len(obstructions[0])
	visited := make(Lab, n_rows)
	for i := range n_rows {
		visited[i] = make([]bool, n_cols)
	}

	visited[guard.row][guard.col] = true

	loop:
	for {
		dcol := 0
		drow := 0
		switch guard.direction {
		case UP:
			drow = -1
			if guard.row + drow < 0 { break loop}
		case DOWN:
			drow = 1
			if guard.row + drow >= n_rows { break loop}
		case LEFT:
			dcol = -1
			if guard.col + dcol < 0 { break loop }
		case RIGHT:
			dcol = 1
			if guard.col + dcol >= n_cols { break loop }
		}

		if obstructions[guard.row + drow][guard.col + dcol] {
			guard.direction = nextDirection(guard.direction)
		} else {
			guard.row = guard.row + drow
			guard.col = guard.col + dcol
			visited[guard.row][guard.col] = true
		}
	}

	return visited
}


func Day06Part1(obstructions Lab, guard Guard) int {
	visited := visitedFields(obstructions, guard)
	return countTrue(visited)
}



func isLoop(obsRow int, obsCol int, obstructions Lab, guard Guard) bool {
	n_rows := len(obstructions)
	n_cols := len(obstructions[0])

	visited := make(map[Guard]bool)
	visited[guard] = true

	for {
		newCol := guard.col
		newRow := guard.row

		switch guard.direction {
		case UP:
			newRow  -= 1
			if newRow < 0 {
				return false 
			}
		case DOWN:
			newRow += 1
			if newRow >= n_rows {
				return false 
			}
		case LEFT:
			newCol -= 1
			if newCol < 0 {
				return false 
			}
		case RIGHT:
			newCol += 1
			if newCol >= n_cols {
				return false 
			}
		}

		if obstructions[newRow][newCol] || (newCol == obsCol && newRow == obsRow) {
			guard.direction = nextDirection(guard.direction)
		} else {
			guard.row = newRow
			guard.col = newCol
			if visited[guard] {
				return true
			}
			visited[guard] = true
		}
	}
}


func Day06Part2(obstructions Lab, guard Guard) int {
	n_rows := len(obstructions)
	n_cols := len(obstructions[0])
	visted := visitedFields(obstructions, guard)

	n := 0
	for row := range n_rows {
		for col := range n_cols {
			if !obstructions[row][col] && visted[row][col] {
				if isLoop(row, col, obstructions, guard) {
					n += 1
				}
			}
		}
	}

	return n
}

func Day06(input string) error {
	lab, guard := Day06ParseInput(input)

	start := time.Now()
	solution1 := Day06Part1(lab, guard)
	stop := time.Now()
	log.Printf("Part 1: %d in %d μs\n", solution1, stop.Sub(start).Microseconds())

	start = time.Now()
	solution2 := Day06Part2(lab, guard)
	stop = time.Now()
	log.Printf("Part 2: %d in %d μs\n", solution2, stop.Sub(start).Microseconds())
	return nil
}

func init() {
	AddSolution(6, Day06)
}
