package main

import (
	"fmt"
	"math"
	"strconv"

	aoc_utils "github.com/kruspe/aoc-utils"
)

func main() {
	input := aoc_utils.Example()
	data := aoc_utils.Data()

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

func solve1(input []string) int {
	number := 50
	zeroCount := 0
	for _, line := range input {
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			number = (number - (rotation % 100) + 100) % 100
		} else if line[0] == 'R' {
			number = (number + rotation) % 100
		}

		if number == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func solve2(input []string) int {
	number := 50
	zeroCount := 0
	for _, line := range input {
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			zeroCount += int(math.Floor(float64(rotation) / 100))
			newNumber := number - (rotation % 100)

			if number != 0 && newNumber < 0 {
				zeroCount++
			}
			number = (newNumber + 100) % 100
			if number == 0 {
				zeroCount++
			}
		} else if line[0] == 'R' {
			newNumber := number + rotation
			zeroCount += int(math.Floor(float64(newNumber) / 100))
			number = newNumber % 100
		}

	}

	return zeroCount
}
