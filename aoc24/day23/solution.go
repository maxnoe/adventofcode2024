package day23

import (
	"fmt"
	"slices"
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

type Network map[string]struct{}

func allConnected(node string, network Network, connections Connections) bool {
	for other := range network {
		if _, found := connections[other][node]; !found {
			return false
		}
	}
	return true
}

func password(network Network) string {
	names := make([]string, len(network))
	i := 0
	for k := range network {
		names[i] = k
		i += 1
	}
	slices.Sort(names)
	return strings.Join(names, ",")
}

func ExpandNetwork(start string, network Network, connections Connections, networks map[string]Network) {

	pwd := password(network)
	if _, found := networks[pwd]; found {
		return
	}

	neighbors, _ := connections[start]
	for n := range neighbors {
		// already in network
		if _, found := network[n]; found {
			continue
		}

		// check if node is connected to all currently connected nodes
		if allConnected(n, network, connections) {
			ExpandNetwork(start, aoc24.CopyAdd(network, n, struct{}{}), connections, networks)
		}
	}
	networks[password(network)] = network
}

func Part2(connections Connections) (string, error) {
	networks := make(map[string]Network)
	for node := range connections {
		ExpandNetwork(node, Network{node: {}}, connections, networks)
	}

	var largest string
	maxLength := 0

	for pwd, network := range networks {
		if len(network) > maxLength {
			maxLength = len(network)
			largest = pwd
		}

	}
	return largest, nil
}

func init() {
	aoc24.AddSolution(23, Parse, Part1, Part2)
}
