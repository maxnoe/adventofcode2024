package day13

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

var test_input = `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

func TestCost(t *testing.T) {
	test_games, _ := Parse(test_input)
	aoc24.AssertEqual(t, len(test_games), 4)

	cost, ok := test_games[0].Cost()
	aoc24.AssertEqual(t, ok, true)
	aoc24.AssertEqual(t, cost, 280)

	cost, ok = test_games[1].Cost()
	aoc24.AssertEqual(t, ok, false)
	aoc24.AssertEqual(t, cost, 0)

	cost, ok = test_games[2].Cost()
	aoc24.AssertEqual(t, ok, true)
	aoc24.AssertEqual(t, cost, 200)

	cost, ok = test_games[3].Cost()
	aoc24.AssertEqual(t, ok, false)
	aoc24.AssertEqual(t, cost, 0)
}
