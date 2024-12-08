package main

import (
	"aoc/internal/grid"
	"aoc/internal/set"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type Pair struct {
	A grid.Point
	B grid.Point
}

func getPairs(points []grid.Point) set.Set[Pair] {
	pairs := make(set.Set[Pair])
	for _, p1 := range points {
		for _, p2 := range points {
			if p1.X < p2.X {
				pairs.Add(Pair{p1, p2})
			} else if p2.X < p1.X {
				pairs.Add(Pair{p2, p1})
			} else if p1.Y < p2.Y {
				pairs.Add(Pair{p1, p2})
			} else if p1 != p2 {
				pairs.Add(Pair{p2, p1})
			}
		}
	}
	return pairs
}

func partOne(myGrid grid.Grid, freqs set.Set[string]) int {
	antinodes := make(set.Set[grid.Point])
	for char := range freqs {
		points := myGrid.AllInstances(char)[char]
		pairs := getPairs(points)
		for pair := range pairs {
			a, b := pair.A, pair.B
			xDiff, yDiff := b.X-a.X, b.Y-a.Y
			an1 := grid.Point{X: a.X - xDiff, Y: a.Y - yDiff}
			an2 := grid.Point{X: b.X + xDiff, Y: b.Y + yDiff}
			if myGrid[an1] != "" {
				antinodes.Add(an1)
			}
			if myGrid[an2] != "" {
				antinodes.Add(an2)
			}
		}
	}
	return len(antinodes)
}

func partTwo() int {
	return 0
}

func main() {
	myGrid := grid.New(input)
	freqs := make(set.Set[string])
	for _, char := range myGrid {
		if char != "." {
			freqs.Add(char)
		}
	}
	fmt.Println("Part 1: ", partOne(myGrid, freqs))
	fmt.Println("Part 2: ", partTwo())
}
