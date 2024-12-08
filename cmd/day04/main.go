package main

import (
	"aoc/internal/grid"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func partOne(myGrid grid.Grid, xPoints []grid.Point) int {
	sum := 0
	for _, xPoint := range xPoints {
		for _, dir := range grid.Dirs() {
			phrase := []string{myGrid[xPoint]}
			nextPoint, nextChar := xPoint, ""
			for i := 0; i < 3; i++ {
				nextPoint, nextChar = myGrid.Next(nextPoint, dir)
				phrase = append(phrase, nextChar)
			}
			if strings.Join(phrase, "") == "XMAS" {
				sum += 1
			}
		}
	}
	return sum
}

func partTwo(myGrid grid.Grid, aPoints []grid.Point) int {
	sum := 0
	for _, aPoint := range aPoints {
		_, nw := myGrid.Next(aPoint, grid.Point{X: -1, Y: -1})
		_, ne := myGrid.Next(aPoint, grid.Point{X: 1, Y: -1})
		_, sw := myGrid.Next(aPoint, grid.Point{X: -1, Y: 1})
		_, se := myGrid.Next(aPoint, grid.Point{X: 1, Y: 1})
		leg1, leg2 := nw+se, ne+sw
		if (leg1 == "MS" || leg1 == "SM") && (leg2 == "MS" || leg2 == "SM") {
			sum += 1
		}
	}
	return sum
}

func main() {
	myGrid := grid.New(input)
	instances := myGrid.AllInstances("X", "A")
	xPoints, aPoints := instances["X"], instances["A"]
	fmt.Println("Part 1: ", partOne(myGrid, xPoints))
	fmt.Println("Part 2: ", partTwo(myGrid, aPoints))
}
