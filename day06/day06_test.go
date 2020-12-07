package daySix

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartOneWithExample(t *testing.T) {
	var list = utils.ReadInputFile("example.txt")

	result := PartOne(list)
	expected := 11
	if result != expected {
		t.Errorf("Day 6, Part 1 (example): %d; want %d", result, expected)
	}
}

func TestPartOne(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartOne(list)
	expected := 6878
	if result != expected {
		t.Errorf("Day 6, Part 1: %d; want %d", result, expected)
	}
}

func TestPartTwoWithExample(t *testing.T) {
	var list = utils.ReadInputFile("example.txt")

	result := PartTwo(list)
	expected := 6
	if result != expected {
		t.Errorf("Day 6, Part 2 (example): %d; want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartTwo(list)
	expected := 3464
	if result != expected {
		t.Errorf("Day 6, Part 2: %d; want %d", result, expected)
	}
}
