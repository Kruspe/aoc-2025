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

type foodRange struct {
	lowerBound int
	upperBound int
}

func solve1(input []string) int {
	separator := aoc_utils.GetSeparators(input)[0]

	freshFoodRanges := make([]*foodRange, 0)
	for _, line := range input[:separator] {
		idRange := strings.Split(line, "-")
		lowerBound, err := strconv.Atoi(idRange[0])
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(idRange[1])
		if err != nil {
			panic(err)
		}
		rangeExtended := false
		for _, freshFoodRange := range freshFoodRanges {
			if lowerBound >= freshFoodRange.lowerBound && lowerBound <= freshFoodRange.upperBound && upperBound >= freshFoodRange.upperBound {
				freshFoodRange.upperBound = upperBound
				rangeExtended = true
				break
			}
			if lowerBound <= freshFoodRange.lowerBound && upperBound >= freshFoodRange.upperBound {
				freshFoodRange.lowerBound = lowerBound
				freshFoodRange.upperBound = upperBound
				rangeExtended = true
				break
			}
			if upperBound <= freshFoodRange.upperBound && upperBound >= freshFoodRange.lowerBound && lowerBound <= freshFoodRange.lowerBound {
				freshFoodRange.lowerBound = lowerBound
				rangeExtended = true
				break
			}
		}
		if !rangeExtended {
			freshFoodRanges = append(freshFoodRanges, &foodRange{
				lowerBound: lowerBound,
				upperBound: upperBound,
			})
		}
	}

	result := 0
	for _, food := range input[separator+1:] {
		id, err := strconv.Atoi(food)
		if err != nil {
			panic(err)
		}
		for _, freshFoodRange := range freshFoodRanges {
			if id >= freshFoodRange.lowerBound && id <= freshFoodRange.upperBound {
				result++
				break
			}
		}
	}

	return result
}

func solve2(input []string) int {
	separator := aoc_utils.GetSeparators(input)[0]

	freshFoodRanges := make([]*foodRange, 0)
	for _, line := range input[:separator] {
		idRange := strings.Split(line, "-")
		lowerBound, err := strconv.Atoi(idRange[0])
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(idRange[1])
		if err != nil {
			panic(err)
		}
		rangeExtended := false
		for _, freshFoodRange := range freshFoodRanges {
			if lowerBound >= freshFoodRange.lowerBound && lowerBound <= freshFoodRange.upperBound && upperBound >= freshFoodRange.upperBound {
				freshFoodRange.upperBound = upperBound
				rangeExtended = true
				break
			}
			if lowerBound <= freshFoodRange.lowerBound && upperBound >= freshFoodRange.upperBound {
				freshFoodRange.lowerBound = lowerBound
				freshFoodRange.upperBound = upperBound
				rangeExtended = true
				break
			}
			if upperBound <= freshFoodRange.upperBound && upperBound >= freshFoodRange.lowerBound && lowerBound <= freshFoodRange.lowerBound {
				freshFoodRange.lowerBound = lowerBound
				rangeExtended = true
				break
			}
		}
		if !rangeExtended {
			freshFoodRanges = append(freshFoodRanges, &foodRange{
				lowerBound: lowerBound,
				upperBound: upperBound,
			})
		}
	}

	for {
		merged := false
		for index, freshFoodRange := range freshFoodRanges[:len(freshFoodRanges)-1] {
			for _, compareRange := range freshFoodRanges[index+1:] {
				if freshFoodRange.lowerBound >= compareRange.lowerBound && freshFoodRange.lowerBound <= compareRange.upperBound && freshFoodRange.upperBound >= compareRange.upperBound {
					compareRange.upperBound = freshFoodRange.upperBound
					merged = true
				}
				if freshFoodRange.lowerBound <= compareRange.lowerBound && freshFoodRange.upperBound >= compareRange.upperBound {
					compareRange.lowerBound = freshFoodRange.lowerBound
					compareRange.upperBound = freshFoodRange.upperBound
					merged = true
				}
				if freshFoodRange.upperBound <= compareRange.upperBound && freshFoodRange.upperBound >= compareRange.lowerBound && freshFoodRange.lowerBound <= compareRange.lowerBound {
					compareRange.lowerBound = freshFoodRange.lowerBound
					merged = true
				}
			}
			if merged == true {
				freshFoodRanges = append(freshFoodRanges[:index], freshFoodRanges[index+1:]...)
				break
			}
		}
		if !merged {
			break
		}
	}

	result := 0
	for _, freshFoodRange := range freshFoodRanges {
		result += freshFoodRange.upperBound - freshFoodRange.lowerBound + 1
	}

	return result
}
