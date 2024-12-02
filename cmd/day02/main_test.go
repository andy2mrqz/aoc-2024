package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input_test.txt
var input_test string

func TestPartTwo(t *testing.T) {
	reports := strings.Split(input_test, "\n")
	expected := 5
	if numValid := partTwo(reports); numValid != expected {
		t.Fatalf("Expected %v valid reports, got %v", expected, numValid)
	}
}
