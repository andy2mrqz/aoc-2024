package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Grid map[Point]string
type Point struct{ x, y int }

func (g Grid) Next(p Point, dir Point) (Point, string) {
	nextPoint := Point{p.x + dir.x, p.y + dir.y}
	nextChar := g[nextPoint]
	return nextPoint, nextChar
}
func partOne(i string) int {
	// create a grid of cells
	grid := Grid{}
	// find all instances of cell 'X'
	xPoints := []Point{}
	for y, line := range strings.Split(i, "\n") {
		for x, char := range line {
			point := Point{x, y}
			if char == 'X' {
				xPoints = append(xPoints, point)
			}
			grid[point] = string(char)
		}
	}
	// perform the search
	sum := 0
	dirs := []Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, xPoint := range xPoints {
		for _, dir := range dirs {
			phrase := []string{grid[xPoint]}
			nextPoint, nextChar := xPoint, ""
			for i := 0; i < 3; i++ {
				nextPoint, nextChar = grid.Next(nextPoint, dir)
				phrase = append(phrase, nextChar)
			}
			if strings.Join(phrase, "") == "XMAS" {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1: ", partOne(input))
	// fmt.Println("Part 2: ", 0)
}
