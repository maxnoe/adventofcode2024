package day09

import (
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


func Parse(input string) ([]int8, error) {
	input = strings.TrimSpace(input)
	disk_space := make([]int8, len(input))

	for i, chr := range input {
		disk_space[i] = int8(chr - '0')
	}

	return disk_space, nil
}


func sum(numbers []int8) int {
	s := 0
	for _, n := range numbers {
		s += int(n)
	}
	return s
}

func CheckSum(blocks []int) int {
	res := 0
	for i, file_id := range blocks {
		if file_id != -1 {
			res += i * file_id
		}
	}

	return res

}

func Compact(blocks []int) {
	space_ptr := 0
	file_ptr := len(blocks) - 1

	for file_ptr > space_ptr {
		for blocks[space_ptr] != -1 {
			space_ptr += 1
		}
		for blocks[file_ptr] == -1 {
			file_ptr -= 1
		}
		if space_ptr >= file_ptr {
			break
		}

		blocks[space_ptr] = blocks[file_ptr]
		blocks[file_ptr] = -1
		space_ptr += 1
		file_ptr -= 1
	}
}


func Part1(disk_space []int8) (int, error) {
	n_blocks := sum(disk_space)
	blocks := make([]int, n_blocks)
	log.Printf("n_blocks = %d", n_blocks)
	
	block := 0
	file_id := 0
	for i, size := range disk_space {

		fill := file_id
		if i % 2 == 1 {
			fill = -1	
		}

		for j := range size {
			blocks[block + int(j)] = fill
		}

		if i % 2 == 0 {
			file_id += 1
		}
		block += int(size)
	}

	Compact(blocks)
	return CheckSum(blocks), nil
}

func Part2(disk_space []int8) (int, error) {
	return 0, nil
}

func init() {
	aoc24.AddSolution(9, Parse, Part1, Part2)
}
