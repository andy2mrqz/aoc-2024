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

func partOne(grid Grid, xPoints []Point) int {
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

func partTwo(grid Grid, aPoints []Point) int {
	sum := 0
	for _, aPoint := range aPoints {
		_, nw := grid.Next(aPoint, Point{-1, -1})
		_, ne := grid.Next(aPoint, Point{1, -1})
		_, sw := grid.Next(aPoint, Point{-1, 1})
		_, se := grid.Next(aPoint, Point{1, 1})
		leg1, leg2 := nw+se, ne+sw
		if (leg1 == "MS" || leg1 == "SM") && (leg2 == "MS" || leg2 == "SM") {
			sum += 1
		}
	}
	return sum
}

func initGrid(i string, charStart1 rune, charStart2 rune) (Grid, []Point, []Point) {
	// create a grid of cells
	grid := Grid{}
	// find all instances of runes of interest
	charPoints1 := []Point{}
	charPoints2 := []Point{}
	for y, line := range strings.Split(i, "\n") {
		for x, char := range line {
			point := Point{x, y}
			if char == charStart1 {
				charPoints1 = append(charPoints1, point)
			} else if char == charStart2 {
				charPoints2 = append(charPoints2, point)
			}
			grid[point] = string(char)
		}
	}
	return grid, charPoints1, charPoints2
}

func main() {
	grid, charPoints1, charPoints2 := initGrid(input, 'X', 'A')
	fmt.Println("Part 1: ", partOne(grid, charPoints1))
	fmt.Println("Part 2: ", partTwo(grid, charPoints2))
}
