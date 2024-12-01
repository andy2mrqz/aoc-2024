package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne(left []int, right []int) int {
        var sum int
        for i := range left {
                diff := left[i] - right[i]
                if diff < 0 {
                        sum -= diff
                } else {
                        sum += diff
                }
        }
        return sum
}

func main() {
        var left, right []int
        for idx, strNum := range strings.Fields(input) {
                number, err := strconv.Atoi(strNum)
                if err != nil {
                        panic(err)
                }
                if idx % 2 == 0 {
                        left = append(left, number)
                } else {
                        right = append(right, number)
                }
        }
        sort.Ints(left)
        sort.Ints(right)

        fmt.Println("Part 1: ", partOne(left, right))
}