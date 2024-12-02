package main

import (
	_ "embed"
	"testing"
)

func TestPartTwo(t *testing.T) {
	valid := [][]int{
		{1, 3, 6, 7, 9},
		{3, 4, 5, 6, 49, 7, 8, 11},
		{49, 3, 4, 5, 6, 7, 8, 11},
		{7, 6, 4, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{49, 48, 46, 46, 43, 42},
		{11, 49, 48, 46, 43, 42},
		{10, 3, 11, 12, 13},
		{10, 15, 9, 8, 7},
		{10, 11, 12, 13, 14, 5, 15, 16, 17},
		{10, 3, 4, 5, 6, 7},
		{10, 14, 15, 16},
		{10, 11, 18, 12},
		{10, 9, 11, 12, 13},
	}
	invalid := [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{49, 3, 4, 5, 49, 6, 7, 8, 11},
		{5, 0, 6, 10},
		{10, 10, 10, 10, 11},
		{1, 4, 7, 11, 15},
	}
	for _, item := range valid {
		if !isValidPartTwo(item) {
			t.Fatalf("Expected a valid report for %v\n", item)
		}
	}
	for _, item := range invalid {
		if isValidPartTwo(item) {
			t.Fatalf("Expected an invalid report for %v\n", item)
		}
	}
}
