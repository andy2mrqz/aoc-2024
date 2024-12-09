package main

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input.txt
var input string

const empty = -1

func partOne(fileBlocks []int) int {
	for leftIdx, rightIdx := 0, len(fileBlocks)-1; leftIdx < rightIdx; {
		if fileBlocks[leftIdx] != empty {
			leftIdx += 1
			continue
		}
		if fileBlocks[rightIdx] == empty {
			rightIdx -= 1
			continue
		}
		fileBlocks[leftIdx] = fileBlocks[rightIdx]
		fileBlocks[rightIdx] = empty
		leftIdx, rightIdx = leftIdx+1, rightIdx-1
	}
	sum := 0
	for i, val := range fileBlocks {
		if val == empty {
			break
		}
		sum += i * val
	}
	return sum
}

func partTwo() int {
	return 0
}

func parseInputToBlocks(i string) []int {
	res := []int{}
	for idx, char := range i {
		if idx%2 == 0 {
			res = append(res, slices.Repeat([]int{idx / 2}, int(char-'0'))...)
		} else {
			res = append(res, slices.Repeat([]int{empty}, int(char-'0'))...)
		}
	}
	return res
}

func main() {
	fileBlocks := parseInputToBlocks(input)
	fmt.Println("Part 1: ", partOne(fileBlocks))
	fmt.Println("Part 2: ", partTwo())
}
