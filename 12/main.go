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

type shape struct {
	requiredSpace int
}

type area struct {
	availableSpace int
	requiredShapes map[string]int
}

func solve1(input []string) int {
	shapes := make(map[string]shape)
	areas := make([]area, 0)
	iterator := 0
	for {
		if iterator == len(input) {
			break
		}

		if !strings.Contains(input[iterator], "x") {
			size := 0
			for i := 1; i < 4; i++ {
				for _, r := range input[iterator+i] {
					if r == '#' {
						size++
					}
				}
			}
			shapes[input[iterator][:len(input[iterator])-1]] = shape{
				requiredSpace: size,
			}

			iterator += 5
		} else {
			split := strings.Split(input[iterator], ": ")
			areaSplit := strings.Split(split[0], "x")
			x, err := strconv.Atoi(areaSplit[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(areaSplit[1])
			if err != nil {
				panic(err)
			}

			requiredShapes := make(map[string]int)
			for i, s := range strings.Split(split[1], " ") {
				amount, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				requiredShapes[strconv.Itoa(i)] = amount
			}
			areas = append(areas, area{
				availableSpace: x * y,
				requiredShapes: requiredShapes,
			})
			iterator++
		}
	}

	result := 0
	for _, a := range areas {
		totalRequiredSpace := 0
		for s, amount := range a.requiredShapes {
			if amount == 0 {
				continue
			}
			totalRequiredSpace += shapes[s].requiredSpace * amount
		}

		if totalRequiredSpace < a.availableSpace {
			result++
		}
	}

	return result
}

func solve2(input []string) int {
	return 0
}
