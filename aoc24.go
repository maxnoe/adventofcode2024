package main

import (
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
)

func main() {
	godotenv.Load()

	args := os.Args

	if len(args) != 2 {
		log.Fatalln("Usage: aoc24 <day>")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Error parsing day as int: %s\n", err)
	}

	log.Printf("Getting input for day %d\n", day)
	input, err := aoc24.GetInput(day)
	if err != nil {
		log.Fatalf("Error getting input for day %d: %s\n", 2, err)
	}
	log.Println("done")

	err = aoc24.SolveDay(day, input);
	if err != nil {
		log.Fatalf("Errror solving day %d: %s", day, err)
	}
}
