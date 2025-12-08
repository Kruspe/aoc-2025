package main

import (
	"fmt"

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

type coordinate struct {
	x int
	y int
}

func solve1(input []string) int {
	var startPoint coordinate
	for x, l := range input[0] {
		if string(l) == "S" {
			startPoint = coordinate{x: x, y: 0}
			break
		}
	}
	beams := map[coordinate]bool{startPoint: true}
	splitCounter := 0
	for {
		if len(beams) == 0 {
			break
		}
		newBeams := make(map[coordinate]bool, len(beams))
		for beam := range beams {
			beam.y++
			if beam.y >= len(input) {
				delete(beams, beam)
				continue
			}
			if string(input[beam.y][beam.x]) == "^" {
				splitCounter++
				beam.x--
				newBeams[coordinate{x: beam.x + 2, y: beam.y}] = true
			}
			newBeams[beam] = true
		}
		beams = newBeams
	}

	return splitCounter
}

func solve2(input []string) int {
	var startPoint coordinate
	for x, l := range input[0] {
		if string(l) == "S" {
			startPoint = coordinate{x: x, y: 0}
			break
		}
	}
	beams := map[coordinate]int{startPoint: 1}
	for {
		newBeams := make(map[coordinate]int)
		for beam, counter := range beams {
			beam.y++
			if beam.y >= len(input) {
				continue
			}
			if string(input[beam.y][beam.x]) == "^" {
				beam.x--
				newBeams[coordinate{x: beam.x + 2, y: beam.y}] += counter
			}
			newBeams[beam] += counter
		}
		if len(newBeams) == 0 {
			break
		}
		beams = newBeams
	}

	result := 0
	for _, count := range beams {
		result += count
	}
	return result
}
