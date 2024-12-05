package main

import (
	"aoc/internal/set"
	"aoc/internal/utils"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func handleRules(beforeSet Rulebook, afterSet Rulebook, updates [][]int) (int, int) {
	part1Sum, part2Sum := 0, 0
	for _, update := range updates {
		isValid := true
		seen := make(set.Set[int])
		shouldNotSeeAgain := make(set.Set[int])
		for currIdx, pageNumber := range update {
			if shouldNotSeeAgain.Has(pageNumber) {
				intersection := beforeSet[pageNumber].Intersection(seen)
				lowestIdx := len(update)
				for _, x := range intersection.Slice() {
					xIdx := slices.Index(update, x)
					lowestIdx = min(xIdx, lowestIdx)
				}
				update = slices.Delete(update, currIdx, currIdx+1)
				update = slices.Insert(update, lowestIdx, pageNumber)
				isValid = false
			}
			afterRules := afterSet[pageNumber]
			shouldNotSeeAgain.Add(afterRules.Slice()...)
			seen.Add(pageNumber)
		}
		middleIdx := len(update) / 2
		if isValid {
			part1Sum += update[middleIdx]
		} else {
			part2Sum += update[middleIdx]
		}
	}
	return part1Sum, part2Sum
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
	part1, part2 := handleRules(beforeSet, afterSet, instructions)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
