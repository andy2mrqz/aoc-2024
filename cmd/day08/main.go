package main

import (
	"aoc/internal/grid"
	"aoc/internal/set"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var input string

type Pair struct {
	A grid.Point
	B grid.Point
}

func getPairs(points []grid.Point) set.Set[Pair] {
	pairs := make(set.Set[Pair])
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			if p1.X > p2.X || (p1.X == p2.X && p1.Y < p2.Y) {
				pairs.Add(Pair{p1, p2})
			} else {
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

type Operator func(a, b float64) float64

var (
	Add Operator = func(a, b float64) float64 { return a + b }
	Sub Operator = func(a, b float64) float64 { return a - b }
)

func partTwo(myGrid grid.Grid, freqs set.Set[string]) int {
	antinodes := make(set.Set[grid.Point])
	for char := range freqs {
		points := myGrid.AllInstances(char)[char]
		pairs := getPairs(points)
		for pair := range pairs {
			antinodes.Add(pair.A, pair.B)
			a, b := pair.A, pair.B
			xDiff, yDiff := b.X-a.X, b.Y-a.Y
			slope := float64(yDiff) / float64(xDiff)
			for _, op := range []Operator{Add, Sub} {
				currX, currY := float64(a.X), float64(a.Y)
				for {
					nextX, nextY := op(currX, 1), op(currY, slope)
					_, fracY := math.Modf(math.Abs(nextY))
					if !(fracY < 1e-9 || fracY > 1.0-1e-9) { // if it's not an integer, continue
						currX, currY = nextX, nextY
						continue
					}
					an := grid.Point{X: int(math.Round(nextX)), Y: int(math.Round(nextY))}
					if myGrid[an] == "" {
						break
					}
					antinodes.Add(an)
					currX, currY = nextX, nextY
				}
			}
		}
	}
	return len(antinodes)
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
	fmt.Println("Part 2: ", partTwo(myGrid, freqs))
}
