package main

import (
	_ "embed"
	"fmt"
	"internal/utils"
	"strings"
)

//go:embed input.txt
var input string

func _getDomSign(ins []int) int {
	sumSigns := 0
	for i, item := range ins[1:] {
		_, sign := utils.AbsInt(item - ins[i])
		sumSigns += sign
	}
	_, domSign := utils.AbsInt(sumSigns)
	return domSign
}

func isValidReport(report []int) bool {
	prev := report[0]
	domSign := _getDomSign(report)
	for _, curr := range report[1:] {
		change, sign := utils.AbsInt(curr - prev)
		isValid := sign == domSign && change >= 1 && change <= 3
		prev = curr
		if !isValid {
			return isValid
		}
	}
	return true
}

func partOne(reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if isValidReport(report) {
			numSafe += 1
		}
	}
	return numSafe
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

func partTwo(reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if isValidPartTwo(report) {
			numSafe += 1
		}
	}
	return numSafe
}

func processInput(i string) [][]int {
	return utils.Map(strings.Split(i, "\n"), func(report string) []int {
		return utils.Map(strings.Fields(report), utils.StrToInt)
	})
}

func main() {
	reports := processInput(input)
	fmt.Println("Part 1: ", partOne(reports))
	fmt.Println("Part 2: ", partTwo(reports))
}
