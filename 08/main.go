package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"

	aoc_utils "github.com/kruspe/aoc-utils"
)

func main() {
	input := aoc_utils.Example()
	data := aoc_utils.Data()

	fmt.Printf("Solution 1: %d\n", solve1(input, 10))
	fmt.Printf("Solution 2: %d\n", solve2(input))

	fmt.Printf("Solution 1: %d\n", solve1(data, 1000))
	fmt.Printf("Solution 2: %d\n", solve2(data))
}

type junction struct {
	x int
	y int
	z int
}

type pair struct {
	j1 junction
	j2 junction
}

type circuit struct {
	junctions []junction
}

func distance(c1, c2 junction) float64 {
	return math.Sqrt(math.Pow(float64(c1.x-c2.x), 2) + math.Pow(float64(c1.y-c2.y), 2) + math.Pow(float64(c1.z-c2.z), 2))
}

type distanceHeap []pair

func (h distanceHeap) Len() int { return len(h) }
func (h distanceHeap) Less(i, j int) bool {
	return distance(h[i].j1, h[i].j2) < distance(h[j].j1, h[j].j2)
}
func (h distanceHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *distanceHeap) Push(x any) {
	*h = append(*h, x.(pair))
}
func (h *distanceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve1(input []string, connections int) int {
	junctions := make([]junction, 0, len(input))
	for _, line := range input {
		s := strings.Split(line, ",")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(s[2])
		if err != nil {
			panic(err)
		}
		junctions = append(junctions, junction{
			x: x,
			y: y,
			z: z,
		})
	}

	h := &distanceHeap{}
	heap.Init(h)
	for index, j1 := range junctions[:len(junctions)-1] {
		for i := index + 1; i < len(junctions); i++ {
			heap.Push(h, pair{j1: j1, j2: junctions[i]})
		}
	}

	circuits := map[junction]*circuit{}
	for i := 0; i < connections; i++ {
		if h.Len() == 0 {
			break
		}
		p := heap.Pop(h).(pair)

		c1, hasJ1 := circuits[p.j1]
		c2, hasJ2 := circuits[p.j2]

		if hasJ1 && hasJ2 && c1 == c2 {
			continue
		}
		if hasJ1 && hasJ2 && c1 != c2 {
			c1.junctions = append(c1.junctions, c2.junctions...)
			for _, j := range c2.junctions {
				circuits[j] = c1
			}
			continue
		}

		if !hasJ1 && !hasJ2 {
			newCircuit := &circuit{junctions: []junction{p.j1, p.j2}}
			circuits[p.j1] = newCircuit
			circuits[p.j2] = newCircuit
		} else if hasJ1 {
			c1.junctions = append(c1.junctions, p.j2)
			circuits[p.j2] = c1
		} else {
			c2.junctions = append(c2.junctions, p.j1)
			circuits[p.j1] = c2
		}
	}

	var a, b, c int
	usedMap := make(map[*circuit]bool)
	for _, circuit := range circuits {
		if usedMap[circuit] {
			continue
		}
		if len(circuit.junctions) > a {
			c, b, a = b, a, len(circuit.junctions)
		} else if len(circuit.junctions) > b {
			c, b = b, len(circuit.junctions)
		} else if len(circuit.junctions) > c {
			c = len(circuit.junctions)
		}

		usedMap[circuit] = true
	}
	return a * b * c
}

func solve2(input []string) int {
	junctions := make([]junction, 0, len(input))
	for _, line := range input {
		s := strings.Split(line, ",")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(s[2])
		if err != nil {
			panic(err)
		}
		junctions = append(junctions, junction{
			x: x,
			y: y,
			z: z,
		})
	}

	h := &distanceHeap{}
	heap.Init(h)
	for index, j1 := range junctions[:len(junctions)-1] {
		for i := index + 1; i < len(junctions); i++ {
			heap.Push(h, pair{j1: j1, j2: junctions[i]})
		}
	}

	circuits := map[junction]*circuit{}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		c1, hasJ1 := circuits[p.j1]
		c2, hasJ2 := circuits[p.j2]

		if hasJ1 && hasJ2 && c1 == c2 {
			continue
		}
		if hasJ1 && hasJ2 && c1 != c2 {
			c1.junctions = append(c1.junctions, c2.junctions...)
			for _, j := range c2.junctions {
				circuits[j] = c1
			}
			continue
		}
		if !hasJ1 && !hasJ2 {
			newCircuit := &circuit{junctions: []junction{p.j1, p.j2}}
			circuits[p.j1] = newCircuit
			circuits[p.j2] = newCircuit
		} else if hasJ1 {
			c1.junctions = append(c1.junctions, p.j2)
			circuits[p.j2] = c1
		} else {
			c2.junctions = append(c2.junctions, p.j1)
			circuits[p.j1] = c2
		}

		if len(circuits[p.j1].junctions) == len(junctions) {
			return p.j1.x * p.j2.x
		}
	}

	var a, b, c int
	usedMap := make(map[*circuit]bool)
	for _, circuit := range circuits {
		if usedMap[circuit] {
			continue
		}
		if len(circuit.junctions) > a {
			c, b, a = b, a, len(circuit.junctions)
		} else if len(circuit.junctions) > b {
			c, b = b, len(circuit.junctions)
		} else if len(circuit.junctions) > c {
			c = len(circuit.junctions)
		}

		usedMap[circuit] = true
	}
	return a * b * c
}
