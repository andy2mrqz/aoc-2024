package grid

import (
	"strings"
)

type Grid map[Point]string
type Point struct{ X, Y int }

func New(i string) Grid {
	grid := Grid{}
	for y, line := range strings.Split(i, "\n") {
		for x, char := range line {
			point := Point{x, y}
			grid[point] = string(char)
		}
	}
	return grid
}

func Dirs() []Point {
	return []Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
}

func (g Grid) Next(p Point, dir Point) (Point, string) {
	nextPoint := Point{p.X + dir.X, p.Y + dir.Y}
	nextChar := g[nextPoint]
	return nextPoint, nextChar
}

func (g Grid) FindNext(searchChar string) (*Point, bool) {
	for gridPoint, gridChar := range g {
		if gridChar == searchChar {
			return &gridPoint, true
		}
	}
	return nil, false
}

func (g Grid) AllInstances(searchChars ...string) map[string][]Point {
	pointsByChar := make(map[string][]Point)
	for gridPoint, gridChar := range g {
		for _, searchChar := range searchChars {
			if gridChar == searchChar {
				pointsByChar[searchChar] = append(pointsByChar[searchChar], gridPoint)
			}
		}
	}
	return pointsByChar
}
