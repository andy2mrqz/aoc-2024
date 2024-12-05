package main

import (
	"aoc/internal/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func isValidReport(report []int) bool {
	prev := report[0]
	_, currSign := utils.AbsInt(report[1] - prev)
	for _, curr := range report[1:] {
		change, sign := utils.AbsInt(curr - prev)
		isValid := sign == currSign && change >= 1 && change <= 3
		prev = curr
		if !isValid {
			return false
		}
	}
	return true
}

func sliceExclude(s []int, idx int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:idx]...)
	return append(ret, s[idx+1:]...)
}

func isValidPartTwo(report []int) bool {
	if isValidReport(report) {
		return true
	}
	for idx := range report {
		if isValidReport(sliceExclude(report, idx)) {
			return true
		}
	}
	return false
}

func processInput(i string) [][]int {
	return utils.Map(strings.Split(i, "\n"), func(report string) []int {
		return utils.Map(strings.Fields(report), utils.StrToInt)
	})
}

func main() {
	reports := processInput(input)
	numSafe, numSafe2 := 0, 0
	for _, report := range reports {
		if isValidReport(report) {
			numSafe += 1
		}
		if isValidPartTwo(report) {
			numSafe2 += 1
		}
	}
	fmt.Println("Part 1: ", numSafe)
	fmt.Println("Part 2: ", numSafe2)
}
