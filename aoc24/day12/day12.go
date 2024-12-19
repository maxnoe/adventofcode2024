package day12

import (
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

func Parse(input string) ([]string, error) {
	return strings.Split(strings.TrimSpace(input), "\n"), nil
}

type Region struct {
	Plant     byte
	Area      int
	Perimeter int
}

func (r *Region) Cost() int {
	return r.Area * r.Perimeter
}

type Pos struct {
	r int
	c int
}

var DIRECTIONS = [4]Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func FindRegions(plots []string) ([]Region, [][]int) {
	n_rows := len(plots)
	n_cols := len(plots[0])

	region_map := make([][]int, n_rows)
	for i := range n_rows {
		region_map[i] = make([]int, n_cols)
	}

	regions := make([]Region, 0)
	current_region_id := 0
	for r := range n_rows {
		for c := range n_cols {
			pos := Pos{r, c}

			if region_map[r][c] != 0 {
				continue
			}

			current_region_id += 1
			current_plant := plots[r][c]
			current_region := Region{current_plant, 0, 0}
			to_check := []Pos{pos}

			for len(to_check) > 0 {
				pos := to_check[0]
				to_check = to_check[1:]
				if region_map[pos.r][pos.c] != 0 {
					continue
				}
				region_map[pos.r][pos.c] = current_region_id
				current_region.Area += 1
				current_region.Perimeter += 4

				for _, d := range DIRECTIONS {
					n := Pos{pos.r + d.r, pos.c + d.c}
					if n.c < 0 || n.c >= n_cols || n.r < 0 || n.r >= n_rows {
						continue
					}
					if plots[n.r][n.c] == current_plant {
						current_region.Perimeter -= 1
						to_check = append(to_check, n)
					}
				}
			}
			regions = append(regions, current_region)
		}
	}

	return regions, region_map
}

func Part1(plots []string) (int, error) {
	cost := 0
	regions, _ := FindRegions(plots)
	for _, region := range regions {
		cost += region.Cost()
	}
	return cost, nil
}

func Part2(plots []string) (int, error) {
	cost := 0
	return cost, nil
}

func init() {
	aoc24.AddSolution(12, Parse, Part1, Part2)
}
