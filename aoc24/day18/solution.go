package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Pos struct {
	X int
	Y int
}

func Parse(input string) ([]Pos, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	bytes := make([]Pos, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Expected 2 elements in line %d, got %d", i, len(parts))
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("Error parsing X as int in line %d: %v", i, err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("Error parsing Y as int in line %d: %v", i, err)
		}
		bytes[i] = Pos{x, y}
	}

	return bytes, nil
}

var directions = []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Head struct {
	pos    Pos
	length int
}

func findShortestPath(cols int, rows int, positions []Pos) int {
	corrupt := make(map[Pos]struct{})
	for _, pos := range positions {
		corrupt[pos] = struct{}{}
	}

	to_check := []Head{{Pos{0, 0}, 1}}
	visited := make(map[Pos]struct{})

	for len(to_check) > 0 {
		head := to_check[0]
		pos := head.pos
		to_check = to_check[1:]

		if _, found := visited[pos]; found {
			continue
		}

		visited[pos] = struct{}{}

		for _, d := range directions {
			n := Pos{pos.X + d.X, pos.Y + d.Y}

			if n.X == cols && n.Y == rows {
				return head.length
			}

			if n.X < 0 || n.X > cols || n.Y < 0 || n.Y > rows {
				continue
			}

			if _, found := corrupt[n]; found {
				continue
			}

			to_check = append(to_check, Head{n, head.length + 1})
		}
	}
	return -1
}

func Part1(positions []Pos) (int, error) {
	return findShortestPath(70, 70, positions[:1024]), nil
}

func part2Impl(cols int, rows int, positions []Pos) (Pos, error) {
	for i, pos := range positions {
		if findShortestPath(cols, rows, positions[:i+1]) == -1 {
			return pos, nil
		}
	}
	return Pos{}, fmt.Errorf("Did not find a byte blocking path")
}

func Part2(positions []Pos) (string, error) {
	pos, err := part2Impl(70, 70, positions)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(pos.X) + "," + strconv.Itoa(pos.Y), nil
}

func init() {
	aoc24.AddSolution(18, Parse, Part1, Part2)
}
