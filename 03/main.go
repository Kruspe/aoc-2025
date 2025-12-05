package main

import (
	"fmt"
	"strconv"
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
	for _, line := range input {
		var largestNumber struct {
			index int
			value int
		}
		for i := 9; i > 0; i-- {
			s := strconv.Itoa(i)
			index := strings.Index(line[:len(line)-1], s)
			if index != -1 {
				largestNumber = struct {
					index int
					value int
				}{index: index, value: i}
				break
			}
		}
		for i := 9; i > 0; i-- {
			s := strconv.Itoa(i)
			index := strings.Index(line[largestNumber.index+1:], s)
			if index != -1 {
				n, err := strconv.Atoi(fmt.Sprintf("%s%s", string(line[largestNumber.index]), s))
				if err != nil {
					panic(err)
				}
				result += n
				break
			}
		}
	}
	return result
}

func solve2(input []string) int {
	result := 0
	for _, line := range input {
		var lineNumbers []int
		for _, s := range line {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				panic(err)
			}
			lineNumbers = append(lineNumbers, n)
		}

		number := ""
		for i := 0; i < 12; i++ {
			largestNumberIndex, largestNumber := 0, 0
			if len(lineNumbers[largestNumberIndex:]) == 12-i {

				number += strings.Trim(strings.Replace(fmt.Sprint(lineNumbers), " ", "", -1), "[]")
				break
			}
			for index, n := range lineNumbers[largestNumberIndex : len(lineNumbers)-(11-i)] {
				if n == 9 {
					largestNumber = n
					largestNumberIndex = index
					break
				}
				if n > largestNumber {
					largestNumber = n
					largestNumberIndex = index
				}
			}
			number += strconv.Itoa(largestNumber)
			newLineNumbers := lineNumbers[largestNumberIndex+1:]
			lineNumbers = newLineNumbers
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		result += n
	}

	return result
}
