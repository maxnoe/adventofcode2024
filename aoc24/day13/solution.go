package day13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Pair struct {
	x int
	y int
}

type Game struct {
	a Pair
	b Pair
	prize Pair
}

func DivOk(nom int, denom int) (int, bool) {
	return nom / denom, nom % denom == 0
}

func (game Game) Cost() (int, bool) {
	nom := game.prize.y * game.a.x - game.prize.x * game.a.y
	denom := game.a.x * game.b.y - game.b.x * game.a.y
	b, ok := DivOk(nom, denom)
	if !ok {
		return 0, false
	}
	a, ok := DivOk(game.prize.x - b * game.b.x, game.a.x)
	if !ok {
		return 0, false
	}

	return 3 * a + b, true
}

var button_re = regexp.MustCompile("Button [AB]: X[+](\\d+), Y[+](\\d+)")
var prize_re = regexp.MustCompile("Prize: X=(\\d+), Y=(\\d+)")


func ParseRe(re *regexp.Regexp, s string) (Pair, error) {
	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return Pair{}, fmt.Errorf("Did not match: %s", s)
	}
	x, err := strconv.Atoi(matches[1])
	if err != nil {
		return Pair{}, err
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		return Pair{}, err
	}
	return Pair{x, y}, nil
}

func ParseGame(s string) (Game, error) {
	lines := strings.Split(s, "\n")
	if len(lines) != 3 {
		return Game{}, fmt.Errorf("Expected 3 lines for game input, got '%s'", s)
	}
	a, err := ParseRe(button_re, lines[0])
	if err != nil {return Game{}, err}
	b, err := ParseRe(button_re, lines[1])
	if err != nil {return Game{}, err}
	prize, err := ParseRe(prize_re, lines[2])
	if err != nil {return Game{}, err}
	return Game{a, b, prize}, nil
}

func Parse(input string) ([]Game, error)  {
	descs := strings.Split(strings.TrimSpace(input), "\n\n")
	games := make([]Game, len(descs))
	for i, desc := range descs {
		game, err := ParseGame(desc)
		if err != nil {return nil, err}
		games[i] = game
	}
	return games, nil
}

func Part1(games []Game) (int, error) {
	n := 0
	for _, game := range games {
		cost, ok := game.Cost()
		if ok {
			n += cost
		}
	}
	return n, nil
}

var offset = 10000000000000

func Part2(games []Game) (int, error) {
	n := 0
	for _, game := range games {
		new_game := Game{game.a, game.b, Pair{offset + game.prize.x, offset + game.prize.y}}
		cost, ok := new_game.Cost()
		if ok {
			n += cost
		}
	}
	return n, nil
}


func init() {
	aoc24.AddSolution(13, Parse, Part1, Part2)
}
