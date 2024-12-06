package main

import (
	"aoc/internal/grid"
	"aoc/internal/set"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

var (
	Up    = grid.Point{X: 0, Y: -1}
	Right = grid.Point{X: 1, Y: 0}
	Down  = grid.Point{X: 0, Y: 1}
	Left  = grid.Point{X: -1, Y: 0}
)

func rotateDir(p grid.Point) grid.Point {
	nextMap := map[grid.Point]grid.Point{
		Up:    Right,
		Right: Down,
		Down:  Left,
		Left:  Up,
	}
	return nextMap[p]
}

func partOne(myGrid grid.Grid, startPoint grid.Point) int {
	currPoint := startPoint
	currDir := Up
	seen := set.New(currPoint)
	for {
		nextPoint, nextChar := myGrid.Next(currPoint, currDir)
		if nextChar == "" {
			break // guard has left the map!
		} else if nextChar == "#" {
			currDir = rotateDir(currDir)
		} else {
			currPoint = nextPoint
		}
		seen.Add(currPoint)
	}
	return len(seen)
}

func partTwo() int {
	return 0
}

func main() {
	myGrid := grid.New(input)
	startPoint, _ := myGrid.FindNext("^")
	fmt.Println("Part 1: ", partOne(myGrid, *startPoint))
	fmt.Println("Part 2: ", partTwo())
}
