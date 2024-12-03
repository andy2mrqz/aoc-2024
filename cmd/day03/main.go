package main

import (
	_ "embed"
	"fmt"
	"internal/utils"
	"regexp"
)

//go:embed input.txt
var input string

func partOne(mults [][]string) int {
	sum := 0
	for _, mult := range mults {
		l, r := mult[1], mult[2]
		sum += utils.StrToInt(l) * utils.StrToInt(r)
	}
	return sum
}

func main() {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	fmt.Println("Part 1: ", partOne(matches))
	fmt.Println("Part 2: ", 0)
}
