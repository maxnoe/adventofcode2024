package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/maxnoe/adventofcode2024/aoc24"
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

	input, err := aoc24.GetInput(day)
	if err != nil {
		log.Fatalf("Error getting input for day %d: %s\n", 2, err);
	}

	switch day {
	case 1: err = aoc24.Day01(input)
	default: log.Fatalf("Day %d not yet solved", day)
	}

	if err != nil {
		log.Fatalf("Errror solving day %d: %s", day, err)
	}
	fmt.Println(input)
}
