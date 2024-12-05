package main

import (
	"aoc/internal/set"
	"aoc/internal/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func partOne(beforeMap Rulebook, afterMap Rulebook, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		isValid := true
		shouldNotSeeAgain := make(set.Set[int])
		for _, pageNumber := range update {
			if shouldNotSeeAgain.Has(pageNumber) {
				isValid = false
				break
			}
			// beforeRules := beforeMap[pageNumber] // n must be before x|y|z
			if afterRules, found := afterMap[pageNumber]; found {
				shouldNotSeeAgain.Add(afterRules.Slice()...)
			}
		}
		if isValid {
			middleIdx := len(update) / 2
			sum += update[middleIdx]
		}
	}
	return sum
}

func partTwo() int {
	sum := 0
	return sum
}

type Rulebook map[int]set.Set[int]

func (r Rulebook) AddOrInit(key int, item int) {
	if _, ok := r[key]; !ok {
		r[key] = make(set.Set[int])
	}
	r[key].Add(item)
}

func processInput(i string) (Rulebook, Rulebook, [][]int) {
	sections := strings.Split(i, "\n\n")
	beforeSet, afterSet := make(Rulebook), make(Rulebook)
	rulesSection := strings.Fields(sections[0])
	for _, rule := range rulesSection {
		pages := strings.Split(rule, "|")
		before, after := utils.StrToInt(pages[0]), utils.StrToInt(pages[1])
		beforeSet.AddOrInit(before, after)
		afterSet.AddOrInit(after, before)
	}
	updates := utils.Map(strings.Fields(sections[1]), func(row string) []int {
		return utils.Map(strings.Split(row, ","), utils.StrToInt)
	})
	return beforeSet, afterSet, updates
}

func main() {
	beforeSet, afterSet, instructions := processInput(input)
	fmt.Println("Part 1: ", partOne(beforeSet, afterSet, instructions))
	fmt.Println("Part 2: ", partTwo())
}
