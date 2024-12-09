package main

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input_test.txt
var input string

const empty = -1

// 0 [0]
// 1 [. .]
// 2 [1 1 1]
// 3 [. . . .]
// 4 [2 2 2 2 2]

// [0 2 2 1 1 1 2 2 2]

func partOne(fileBlocks [][]int) int {
	// res := []int{}
	for leftSliceIdx, rightSliceIdx := 1, len(fileBlocks)-1; leftSliceIdx < rightSliceIdx; {
		leftSlice, rightSlice := fileBlocks[leftSliceIdx], fileBlocks[rightSliceIdx]
		for i, rightVal := range slices.Backward(rightSlice) {
			if rightVal == empty {
				continue
			}
			for j, leftVal := range leftSlice {
				if leftVal != empty {
					continue
				}
				fileBlocks[leftSliceIdx][j] = rightVal
				fileBlocks[rightSliceIdx][i] = empty
				break
			}
		}
		leftSliceIdx += 1
		rightSliceIdx -= 1
	}
	fmt.Println(fileBlocks)
	return 0
}

func partTwo() int {
	return 0
}

func parseInputToBlocks(i string) [][]int {
	res := [][]int{}
	for idx, char := range i {
		if idx%2 == 0 {
			res = append(res, slices.Repeat([]int{idx / 2}, int(char-'0')))
		} else {
			res = append(res, slices.Repeat([]int{empty}, int(char-'0')))
		}
	}
	return res
}

func main() {
	fileBlocks := parseInputToBlocks(input)
	fmt.Println(fileBlocks)
	fmt.Println("Part 1: ", partOne(fileBlocks))
	fmt.Println("Part 2: ", partTwo())
}
