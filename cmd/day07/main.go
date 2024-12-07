package main

import (
	"aoc/internal/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Operator func(a, b int) int

var (
	Add    Operator = func(a, b int) int { return a + b }
	Mul    Operator = func(a, b int) int { return a * b }
	Concat Operator = func(a, b int) int { return utils.StrToInt(strconv.Itoa(a) + strconv.Itoa(b)) }
)

type Equation struct {
	TestValue int
	Numbers   []int
}

// took just this part from here https://stackoverflow.com/a/29023972
func product(n, repeat int) [][]int {
	ix := make([]int, repeat)
	result := [][]int{}
	for {
		result = append(result, append([]int(nil), ix...))
		j := repeat - 1
		for ; j >= 0 && ix[j] == n-1; j-- {
			ix[j] = 0
		}
		if j < 0 {
			return result
		}
		ix[j]++
	}
}

func (e Equation) solveEquation(operators []Operator) (total int, ok bool) {
	goal, nums := e.TestValue, e.Numbers
	cs := product(len(operators), len(nums)-1)
	for _, c := range cs {
		a := nums[0]
		total := 0
		operatorCombo := []Operator{}
		for _, opIdx := range c {
			operatorCombo = append(operatorCombo, operators[opIdx])
		}
		for idx, b := range nums[1:] {
			total = operatorCombo[idx](a, b)
			a = total
		}
		if total == goal {
			return total, true
		}
	}
	return 0, false
}

func solve(equations []Equation, operators []Operator) int {
	sum := 0
	for _, equation := range equations {
		if total, ok := equation.solveEquation(operators); ok {
			sum += total
		}
	}
	return sum
}

func processInput(i string) []Equation {
	equations := []Equation{}
	for _, line := range strings.Split(i, "\n") {
		parts := strings.Split(line, ":")
		testVal, rest := parts[0], parts[1]
		nums := []int{}
		for _, numStr := range strings.Fields(rest) {
			nums = append(nums, utils.StrToInt(numStr))
		}
		equations = append(equations, Equation{TestValue: utils.StrToInt(testVal), Numbers: nums})
	}
	return equations
}

func main() {
	equations := processInput(input)
	part1Operators := []Operator{Add, Mul}
	part2Operators := []Operator{Add, Mul, Concat}
	fmt.Println("Part 1: ", solve(equations, part1Operators))
	fmt.Println("Part 2: ", solve(equations, part2Operators))
}
