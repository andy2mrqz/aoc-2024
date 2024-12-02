package main

import (
	_ "embed"
	"fmt"
	"internal/utils"
	"strings"
)

//go:embed input.txt
var input string

func partOne(reports [][]int) int {
	numSafe := 0
	for _, levels := range reports {
		prev := levels[0]
		isSafe := false
		_, currSign := utils.AbsInt(levels[1] - prev)
		for _, level := range levels[1:] {
			change, sign := utils.AbsInt(level - prev)
			isSafe = sign == currSign && sign != 0 && change >= 1 && change <= 3
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

func _getDomSign(ins []int) int {
	sumSigns := 0
	for i, item := range ins[1:] {
		_, sign := utils.AbsInt(item - ins[i])
		sumSigns += sign
	}
	_, domSign := utils.AbsInt(sumSigns)
	return domSign
}

func peek(levels []int, i int) (int, bool) {
	if i < len(levels) {
		return levels[i], true
	} else {
		return 0, false
	}
}

func isValidStep(prev int, curr int, domSign int) bool {
	change, sign := utils.AbsInt(curr - prev)
	return sign == domSign && change >= 1 && change <= 3
}

func partTwo(reports [][]int) int {
	numSafe := 0
	for _, levels := range reports {
		prev := levels[0]
		numUnsafe := 0
		domSign := _getDomSign(levels)
		for idx, curr := range levels[1:] {
			isValid := isValidStep(prev, curr, domSign)
			if isValid {
				prev = curr
				continue
			}
			numUnsafe += 1
			next, ok := peek(levels, idx+2)
			if ok {
				nextValid := isValidStep(curr, next, domSign)
				if nextValid {
					prev = curr
				}
			}
		}
		if numUnsafe <= 1 {
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
