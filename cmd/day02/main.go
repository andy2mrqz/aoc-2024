package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne(reports []string) int {
	numSafe := 0
	for _, report := range reports {
		levels := strings.Fields(report)
		prev, _ := strconv.Atoi(levels[0])
		isSafe := false
		dir := 0
		for _, strLevel := range levels[1:] {
			level, _ := strconv.Atoi(strLevel)
			change := level - prev
			if dir == 0 {
				dir = change
				if dir == 0 {
					break
				}
			}
			absChange := change
			if change < 0 {
				absChange *= -1
			}
			isSafe = change*dir > 0 && absChange >= 1 && absChange <= 3
			if !isSafe {
				break
			}
			prev = level
		}
		if isSafe {
			numSafe += 1
		}
	}
	return numSafe
}

func partTwo() int {
	return 0
}

func main() {
	reports := strings.Split(input, "\n")
	fmt.Println("Part 1: ", partOne(reports))
	fmt.Println("Part 2: ", partTwo())
}
