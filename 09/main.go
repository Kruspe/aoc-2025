package main

import (
	"fmt"
	"image"
	"math"
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

func getCoordinates(line string) (int, int) {
	split := strings.Split(line, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func solve1(input []string) int {
	var largestArea int
	for index, line := range input {
		x, y := getCoordinates(line)
		for i := index + 1; i < len(input); i++ {
			x1, y1 := getCoordinates(input[i])
			area := int((math.Abs(float64(x-x1)) + 1) * (math.Abs(float64(y-y1)) + 1))
			largestArea = max(largestArea, area)
		}
	}
	return largestArea
}

func solve2(input []string) int {
	points, rectangles, lines := make([]image.Point, 0, len(input)), make([]image.Rectangle, 0), make([]image.Rectangle, 0)
	for _, line := range input {
		x, y := getCoordinates(line)
		p := image.Point{X: x, Y: y}
		for _, p1 := range points {
			r := image.Rectangle{Min: p1, Max: p}.Canon()
			r.Max = r.Max.Add(image.Point{X: 1, Y: 1})
			rectangles = append(rectangles, r)
		}
		points = append(points, p)

		if len(points) > 1 {
			lines = append(lines, image.Rectangle{Min: points[len(points)-2], Max: p}.Canon())
		}
	}
	lines = append(lines, image.Rectangle{Min: points[0], Max: points[len(points)-1]}.Canon())

	largestArea := 0
loop:
	for _, r := range rectangles {
		area := r.Dx() * r.Dy()
		for _, l := range lines {
			l.Max = l.Max.Add(image.Point{X: 1, Y: 1})
			if l.Overlaps(r.Inset(1)) {
				continue loop
			}
		}
		largestArea = max(largestArea, area)
	}
	return largestArea
}
