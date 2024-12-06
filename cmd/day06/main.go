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

func partOne(myGrid grid.Grid, startPoint grid.Point) (numPositions int, isCycle bool, seen set.Set[grid.Point]) {
	currPoint, currDir := startPoint, Up
	seen = set.New(currPoint)
	isCycle = true
	for i := 0; i < len(myGrid); i++ {
		nextPoint, nextChar := myGrid.Next(currPoint, currDir)
		if nextChar == "" {
			isCycle = false
			break // guard has left the map!
		} else if nextChar == "#" {
			currDir = rotateDir(currDir)
		} else {
			currPoint = nextPoint
		}
		seen.Add(currPoint)
	}
	return len(seen), isCycle, seen
}

func partTwo(myGrid grid.Grid, startPoint grid.Point, seen set.Set[grid.Point]) int {
	numCycles := 0
	seen.Remove(startPoint) // can't add an obstacle to the starting point
	for point := range seen {
		char := myGrid[point]
		myGrid[point] = "#" // test adding a boundary
		if _, isCycle, _ := partOne(myGrid, startPoint); isCycle {
			numCycles += 1
		}
		myGrid[point] = char // reset to original char
	}
	return numCycles
}

func main() {
	myGrid := grid.New(input)
	startPoint, _ := myGrid.FindNext("^")
	numPositions, _, seen := partOne(myGrid, *startPoint)
	fmt.Println("Part 1: ", numPositions)
	fmt.Println("Part 2: ", partTwo(myGrid, *startPoint, seen))
}
