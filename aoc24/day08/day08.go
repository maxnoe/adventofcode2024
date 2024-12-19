package day08

import (
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Pos struct {
	Row int
	Col int
}

type World struct {
	Antennas map[rune][]Pos
	Rows     int
	Cols     int
}

func Parse(input string) (World, error) {

	lines := strings.Split(strings.TrimSpace(input), "\n")

	world := World{make(map[rune][]Pos), len(lines), len(lines[0])}

	for i, line := range lines {
		for j, chr := range line {
			if chr != '.' {
				world.Antennas[chr] = append(world.Antennas[chr], Pos{i, j})
			}
		}
	}

	return world, nil
}

func Inside(pos Pos, world World) bool {
	if pos.Col < 0 || pos.Col >= world.Cols {
		return false
	}

	if pos.Row < 0 || pos.Row >= world.Rows {
		return false
	}
	return true
}

func Part1(world World) (int, error) {
	resonances := make(map[Pos]struct{})

	for _, antennas := range world.Antennas {
		for i, a := range antennas {
			for _, b := range antennas[i+1:] {
				drow := b.Row - a.Row
				dcol := b.Col - a.Col

				pos1 := Pos{a.Row - drow, a.Col - dcol}
				if Inside(pos1, world) {
					resonances[pos1] = struct{}{}
				}

				pos2 := Pos{b.Row + drow, b.Col + dcol}
				if Inside(pos2, world) {
					resonances[pos2] = struct{}{}
				}
			}
		}
	}
	return len(resonances), nil
}

func Part2(world World) (int, error) {
	resonances := make(map[Pos]struct{})

	for _, antennas := range world.Antennas {
		for i, a := range antennas {
			for _, b := range antennas[i+1:] {
				drow := b.Row - a.Row
				dcol := b.Col - a.Col
				gcd := aoc24.GCD(drow, dcol)
				drow = drow / gcd
				dcol = dcol / gcd

				pos := a
				for Inside(pos, world) {
					resonances[pos] = struct{}{}
					pos = Pos{pos.Row + drow, pos.Col + dcol}
				}

				pos = a
				for Inside(pos, world) {
					resonances[pos] = struct{}{}
					pos = Pos{pos.Row - drow, pos.Col - dcol}
				}
			}
		}
	}
	return len(resonances), nil
}

func init() {
	aoc24.AddSolution(8, Parse, Part1, Part2)
}
