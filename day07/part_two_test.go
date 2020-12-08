package daySeven

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartTwoWithExample(t *testing.T) {
	var list = utils.ReadInputFile("example.txt")

	result := PartTwo(list)
	expected := 32
	if result != expected {
		t.Errorf("Day 07, Part 2 (example): %d; want %d", result, expected)
	}
}

func TestPartTwoWithExampleTwo(t *testing.T) {
	var list = utils.ReadInputFile("example2.txt")

	result := PartTwo(list)
	expected := 126
	if result != expected {
		t.Errorf("Day 07, Part 2 (example 2): %d; want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartTwo(list)
	expected := 85324
	if result != expected {
		t.Errorf("Day 07, Part 2: %d; want %d", result, expected)
	}
}
