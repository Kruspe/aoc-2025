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

type problem struct {
	values               []int
	lengthOfLongestValue int
	operator             string
}

func solve1(input []string) int {
	problems := make([]problem, 0)
	operators := strings.Split(input[len(input)-1], " ")
	for _, o := range operators {
		if len(o) > 0 {
			problems = append(problems, problem{
				operator: o,
			})
		}
	}

	for _, line := range input[:len(input)-1] {
		problemIterator := 0
		split := strings.Split(line, " ")
		for _, s := range split {
			if len(s) > 0 {
				number, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				problems[problemIterator].values = append(problems[problemIterator].values, number)
				problemIterator++
			}
		}
	}

	result := 0
	for _, p := range problems {
		if p.operator == "+" {
			sum := 0
			for _, value := range p.values {
				sum += value
			}
			result += sum
		} else if p.operator == "*" {
			product := 1
			for _, value := range p.values {
				product *= value
			}
			result += product
		}
	}

	return result
}

func solve2(input []string) int {
	problems := make([]*problem, 0)
	operators := strings.Split(input[len(input)-1], " ")
	for _, o := range operators {
		if len(o) > 0 {
			problems = append(problems, &problem{
				operator: o,
			})
		}
	}
	for _, line := range input {
		problemIterator := 0
		split := strings.Split(line, " ")
		for _, s := range split {
			if len(s) > 0 {
				if problems[problemIterator].lengthOfLongestValue < len(s) {
					problems[problemIterator].lengthOfLongestValue = len(s)
				}
				problemIterator++
			}
		}
	}

	indexCounter := 0
	for _, p := range problems {
		values := make([]string, 0, len(input)-1)
		for _, line := range input[:len(input)-1] {
			v := line[indexCounter : indexCounter+p.lengthOfLongestValue]
			values = append(values, v)
		}
		for i := p.lengthOfLongestValue - 1; i >= 0; i-- {
			var newValue string
			for _, v := range values {
				newValue += string(v[i])
			}
			newValue = strings.TrimSpace(newValue)
			number, err := strconv.Atoi(newValue)
			if err != nil {
				panic(err)
			}
			p.values = append(p.values, number)
		}

		indexCounter += p.lengthOfLongestValue + 1
	}

	result := 0
	for _, p := range problems {
		if p.operator == "+" {
			sum := 0
			for _, value := range p.values {
				sum += value
			}
			result += sum
		} else if p.operator == "*" {
			product := 1
			for _, value := range p.values {
				product *= value
			}
			result += product
		}
	}

	return result
}
