package main

import (
	"fmt"
	"strings"

	aoc_utils "github.com/kruspe/aoc-utils"
)

func main() {
	input := aoc_utils.Example()
	input2 := aoc_utils.Example2()
	data := aoc_utils.Data()

	fmt.Printf("Solution 1: %d\n", solve1(input))
	fmt.Printf("Solution 2: %d\n", solve2(input2))

	fmt.Printf("Solution 1: %d\n", solve1(data))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

type device struct {
	name string
	outs []string
}

func solve1(input []string) int {
	deviceMap := make(map[string]device)
	for _, line := range input {
		split := strings.Split(line, ": ")
		outs := strings.Split(split[1], " ")
		deviceMap[split[0]] = device{
			name: split[0],
			outs: outs,
		}
	}

	return findPath(deviceMap, make(map[string]*int), "you", "out")
}

func findPath(deviceMap map[string]device, memo map[string]*int, node, endNode string) int {
	if memo[node] != nil {
		return *memo[node]
	}

	if node == endNode {
		return 1
	}
	if len(deviceMap[node].outs) == 0 {
		return 0
	}

	count := 0
	for _, d := range deviceMap[node].outs {
		count += findPath(deviceMap, memo, d, endNode)
	}
	memo[node] = &count

	return count
}

func solve2(input []string) int {
	deviceMap := make(map[string]device)
	for _, line := range input {
		split := strings.Split(line, ": ")
		outs := strings.Split(split[1], " ")
		deviceMap[split[0]] = device{
			name: split[0],
			outs: outs,
		}
	}

	return findPath(deviceMap, make(map[string]*int), "svr", "dac")*
		findPath(deviceMap, make(map[string]*int), "dac", "fft")*
		findPath(deviceMap, make(map[string]*int), "fft", "out") +
		findPath(deviceMap, make(map[string]*int), "svr", "fft")*
			findPath(deviceMap, make(map[string]*int), "fft", "dac")*
			findPath(deviceMap, make(map[string]*int), "dac", "out")
}
