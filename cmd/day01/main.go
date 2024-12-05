package main

import (
	"aoc/internal/utils"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func partOne(left []int, right []int) int {
	var sum int
	for i := range left {
		diff, _ := utils.AbsInt(left[i] - right[i])
		sum += diff
	}
	return sum
}

func partTwo(left []int, right []int) int {
	simScore := 0
	cache := make(map[int]int) // number, occurrences
	for _, num := range left {
		occ, ok := cache[num]
		if ok {
			simScore += num * occ
			continue
		}
		count := 0
		for j := range right {
			if num == right[j] {
				count += 1
			}
		}
		simScore += num * count
		cache[num] = count
	}
	return simScore
}

func main() {
	var left, right []int
	for idx, strNum := range strings.Fields(input) {
		number := utils.StrToInt(strNum)
		if idx%2 == 0 {
			left = append(left, number)
		} else {
			right = append(right, number)
		}
	}
	slices.Sort(left)
	slices.Sort(right)

	fmt.Println("Part 1: ", partOne(left, right))
	fmt.Println("Part 2: ", partTwo(left, right))
}
