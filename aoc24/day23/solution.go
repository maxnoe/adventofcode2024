package day23

import (
	"fmt"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


type Connections map[string]map[string]struct{}

func Parse(input string) (Connections, error) {
	connections := make(Connections)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		names := strings.Split(line, "-")
		if len(names) != 2 {
			return nil, fmt.Errorf("Expected line of form <name>-<name>, got %s", line)
		}

		for _, k := range names {
			if _, ok := connections[k]; !ok {
				connections[k] = make(map[string]struct{})
			}
		}

		a := names[0]
		b := names[1]
		connections[a][b] = struct{}{}
		connections[b][a] = struct{}{}
	}

	return connections, nil
}


func Part1(connections Connections) (int, error) {
	networks := 0

	for node, neighbors := range connections {
		for neighbor := range neighbors {
			for third := range connections[neighbor] {
				_, found := connections[third][node]
				if found && (node[0] == 't' || neighbor[0] == 't' || third[0] == 't') {
					networks += 1
				}
			}
		}
	} 


	return networks / 6, nil
}

func Part2(Connections Connections) (int, error) {
	return 0, nil
}


func init() {
	aoc24.AddSolution(23, Parse, Part1, Part2)
}
