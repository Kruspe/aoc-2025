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
	invalidIds := make([]int, 0)
	idRanges := strings.Split(input[0], ",")
	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")
		id1, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		id2, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}
		for i := id1; i <= id2; i++ {
			stringId := strconv.Itoa(i)
			if stringId[:len(stringId)/2] == stringId[len(stringId)/2:] {
				invalidIds = append(invalidIds, i)
			}
		}
	}
	result := 0
	for _, id := range invalidIds {
		result += id
	}
	return result
}

func solve2(input []string) int {
	invalidIds := make([]int, 0)
	idRanges := strings.Split(input[0], ",")
	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")
		id1, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		id2, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}
		for i := id1; i <= id2; i++ {
			stringId := strconv.Itoa(i)
			if stringId[:len(stringId)/2] == stringId[len(stringId)/2:] {
				invalidIds = append(invalidIds, i)
				continue
			}
			for j := 0; j < len(stringId)/2; j++ {
				pattern := strings.Repeat(stringId[:j+1], (len(stringId)-j-1)/(j+1))
				if pattern == stringId[j+1:] {
					invalidIds = append(invalidIds, i)
					break
				}
			}
		}
	}
	result := 0
	for _, id := range invalidIds {
		result += id
	}
	return result
}
