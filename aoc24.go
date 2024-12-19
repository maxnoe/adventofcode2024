package main

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/maxnoe/adventofcode2024/aoc24"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day01"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day02"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day03"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day04"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day06"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day07"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day08"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day09"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day10"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day11"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day12"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day13"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day16"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day17"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day18"
	_ "github.com/maxnoe/adventofcode2024/aoc24/day19"
)

func main() {
	godotenv.Load()

	args := os.Args

	if len(args) == 1 || len(args) > 3 {
		log.Fatalln("Usage: aoc24 <day> [input]")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Error parsing day as int: %s\n", err)
	}

	var input string
	if len(args) == 2 {
		log.Printf("Getting input for day %d\n", day)
		input, err = aoc24.GetInput(day)
		if err != nil {
			log.Fatalf("Error getting input for day %d: %s\n", day, err)
		}
		log.Println("done")
	} else {
		var data []byte
		if args[2] == "-" {
			log.Println("Reading input from stdin")
			data, err = io.ReadAll(os.Stdin)
		} else {
			log.Printf("Reading input from %s", args[2])
			data, err = os.ReadFile(args[2])
		}
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		input = string(data)
	}


	err = aoc24.SolveDay(day, input);
	if err != nil {
		log.Fatalf("Errror solving day %d: %s", day, err)
	}
}
