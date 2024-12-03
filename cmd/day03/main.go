package main

import (
	_ "embed"
	"fmt"
	"internal/utils"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func partOne(i string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	mults := re.FindAllStringSubmatch(i, -1)
	sum := 0
	for _, mult := range mults {
		l, r := utils.StrToInt(mult[1]), utils.StrToInt(mult[2])
		sum += l * r
	}
	return sum
}

func partTwo(i string) int {
	doRe := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	insts := doRe.FindAllStringSubmatch(i, -1)
	active := true
	sum := 0
	for _, inst := range insts {
		if strings.HasPrefix(inst[0], "do") {
			active = inst[0] == "do()" // update active according to do() or don't()
		}
		if active && strings.HasPrefix(inst[0], "mul") {
			l, r := utils.StrToInt(inst[1]), utils.StrToInt(inst[2])
			sum += l * r
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1: ", partOne(input))
	fmt.Println("Part 2: ", partTwo(input))
}
