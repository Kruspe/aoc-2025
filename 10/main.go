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

type button struct {
	index    int
	switches map[int]bool
}

type machine struct {
	lights   map[int]bool
	buttons  []button
	joltages []int
}

func getMachine(line string) machine {
	m := machine{
		lights: make(map[int]bool),
	}
	joltageStart := strings.Index(line, "{") + 1
	for _, j := range strings.Split(line[joltageStart:len(line)-1], ",") {
		n, err := strconv.Atoi(j)
		if err != nil {
			panic(err)
		}
		m.joltages = append(m.joltages, n)
	}
	lightEnd := strings.Index(line, "]")
	for index, l := range strings.Split(line[1:lightEnd], "") {
		if l == "." {
			m.lights[index] = false
		} else {
			m.lights[index] = true
		}
	}
	for index, b := range strings.Split(line[lightEnd+2:joltageStart-2], " ") {
		switchMap := make(map[int]bool)
		for _, lightIndex := range strings.Split(b[1:len(b)-1], ",") {
			l, err := strconv.Atoi(lightIndex)
			if err != nil {
				panic(err)
			}
			switchMap[l] = true
		}
		m.buttons = append(m.buttons, button{
			index:    index,
			switches: switchMap,
		})
	}
	return m
}

func compareLightStatus(a, b map[int]bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func solve1(input []string) int {
	result := 0
	for _, line := range input {
		m := getMachine(line)
		combs := aoc_utils.Combinations(m.buttons)
		fewestPresses := 9999999
		for _, comb := range combs {
			lightStatus := make(map[int]bool)
			counter := 0
			for _, button := range comb {
				if counter >= fewestPresses {
					break
				}
				for lightIndex := range button.switches {
					lightStatus[lightIndex] = !lightStatus[lightIndex]
				}
				counter++
				if compareLightStatus(m.lights, lightStatus) && counter < fewestPresses {
					fewestPresses = counter
					break
				}
			}
		}
		result += fewestPresses
	}
	return result
}

type problem struct {
	combs         [][]button
	costs         map[string]*pattern
	joltageLength int
	buttonLength  int
	cache         map[string]*int
}

func createProblem(m machine) problem {
	combs := append(aoc_utils.Combinations(m.buttons), []button{})
	zero := 0
	return problem{
		combs:         combs,
		costs:         getCosts(combs, len(m.joltages), len(m.buttons)),
		joltageLength: len(m.joltages),
		buttonLength:  len(m.buttons),
		cache:         map[string]*int{strings.Repeat("0", len(m.joltages)): &zero},
	}
}

func (p problem) getLowestPresses(joltages []int) int {
	joltageIdentifier := ""
	for _, j := range joltages {
		joltageIdentifier += strconv.Itoa(j)
	}
	if p.cache[joltageIdentifier] != nil {
		return *p.cache[joltageIdentifier]
	}

	repr := getJoltageOnOffRepresentation(joltages)
	validCombs := getValidCombs(repr, p.combs)
	minPresses := 999999
loop:
	for _, comb := range validCombs {
		identifier := []rune(strings.Repeat("0", p.buttonLength))
		for _, b := range comb {
			identifier[b.index] = '1'
		}
		nextJoltage := make([]int, p.joltageLength)
		copy(nextJoltage, joltages)
		for k, v := range p.costs[string(identifier)].value {
			nextJoltage[k] -= v
			if nextJoltage[k] < 0 {
				continue loop
			}
		}
		nextJoltage = getNextJoltage(nextJoltage)
		if p.cache[joltageIdentifier] != nil {
			return *p.cache[joltageIdentifier]
		}
		y := 2*p.getLowestPresses(nextJoltage) + len(comb)
		minPresses = min(minPresses, y)
	}
	p.cache[joltageIdentifier] = &minPresses

	return minPresses
}

func solve2(input []string) int {
	result := 0
	for _, line := range input {
		m := getMachine(line)
		p := createProblem(m)

		presses := p.getLowestPresses(m.joltages)
		result += presses
	}
	return result
}

type pattern struct {
	value map[int]int
	cost  int
}

func getCosts(combs [][]button, patternLength, buttonLength int) map[string]*pattern {
	result := make(map[string]*pattern)
	for _, comb := range combs {
		identifier := []rune(strings.Repeat("0", buttonLength))
		value := make(map[int]int, patternLength)
		for _, b := range comb {
			identifier[b.index] = '1'
			for lightIndex := range b.switches {
				value[lightIndex]++
			}
		}
		result[string(identifier)] = &pattern{
			value: value,
			cost:  len(comb),
		}
	}
	return result
}

func getJoltageOnOffRepresentation(joltage []int) map[int]bool {
	result := make(map[int]bool)
	for index, j := range joltage {
		result[index] = j%2 != 0
	}
	return result
}

func getValidCombs(target map[int]bool, combs [][]button) [][]button {
	var result [][]button
	for _, comb := range combs {
		s := make(map[int]bool)
		for _, b := range comb {
			for lightIndex := range b.switches {
				s[lightIndex] = !s[lightIndex]
			}
		}
		if compareLightStatus(target, s) {
			result = append(result, comb)
		}
	}
	return result
}

func getNextJoltage(joltages []int) []int {
	nextRepr := make([]int, len(joltages))
	for i, joltage := range joltages {
		if joltage%2 != 0 {
			nextRepr[i] = (joltage - 1) / 2
		} else {
			nextRepr[i] = joltage / 2
		}
	}
	return nextRepr
}
