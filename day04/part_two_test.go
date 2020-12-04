package dayFour

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartTwoWithExample(t *testing.T) {
	var list = utils.ReadInputFile("example.txt")

	result := PartTwo(list)
	expected := 2
	if result != expected {
		t.Errorf("Day 4, Part 2 (example): %d; want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartTwo(list)
	expected := 121
	if result != expected {
		t.Errorf("Day 4, Part 2: %d; want %d", result, expected)
	}
}
