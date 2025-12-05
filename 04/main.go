package main

import (
	"fmt"
	"strings"

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
	result := 0
	var grid [][]string
	for _, line := range input {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	paddedGrid := aoc_utils.Pad2dStringArray(grid, ".")

	for y, row := range paddedGrid[1 : len(paddedGrid)-1] {
		for x, s := range row[1 : len(row)-1] {
			if s == "." {
				continue
			}

			counter := 0
			if paddedGrid[y][x] == "@" {
				counter++
			}
			if paddedGrid[y][x+1] == "@" {
				counter++
			}
			if paddedGrid[y][x+2] == "@" {
				counter++
			}
			if paddedGrid[y+1][x] == "@" {
				counter++
			}
			if paddedGrid[y+1][x+2] == "@" {
				counter++
			}
			if paddedGrid[y+2][x] == "@" {
				counter++
			}
			if paddedGrid[y+2][x+1] == "@" {
				counter++
			}
			if paddedGrid[y+2][x+2] == "@" {
				counter++
			}
			if counter < 4 {
				result++
			}
		}
	}

	return result
}

func solve2(input []string) int {
	result := 0
	var grid [][]string
	for _, line := range input {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	paddedGrid := aoc_utils.Pad2dStringArray(grid, ".")

	for {
		var rollsToRemove []struct{ x, y int }
		for y, row := range paddedGrid[1 : len(paddedGrid)-1] {
			for x, s := range row[1 : len(row)-1] {
				if s == "." {
					continue
				}

				counter := 0
				if paddedGrid[y][x] == "@" {
					counter++
				}
				if paddedGrid[y][x+1] == "@" {
					counter++
				}
				if paddedGrid[y][x+2] == "@" {
					counter++
				}
				if paddedGrid[y+1][x] == "@" {
					counter++
				}
				if paddedGrid[y+1][x+2] == "@" {
					counter++
				}
				if paddedGrid[y+2][x] == "@" {
					counter++
				}
				if paddedGrid[y+2][x+1] == "@" {
					counter++
				}
				if paddedGrid[y+2][x+2] == "@" {
					counter++
				}
				if counter < 4 {
					rollsToRemove = append(rollsToRemove, struct{ x, y int }{x: x + 1, y: y + 1})
				}
			}
		}
		if len(rollsToRemove) == 0 {
			break
		}
		result += len(rollsToRemove)
		for _, roll := range rollsToRemove {
			paddedGrid[roll.y][roll.x] = "."
		}
	}
	return result
}
