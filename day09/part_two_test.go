package dayNine

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartTwoWithExample(t *testing.T) {
	var list = utils.ReadInputFileToInts("example.txt")

	result := PartTwo(list, 5)
	expected := 62
	if result != expected {
		t.Errorf("Day 07, Part 2 (example): %d; want %d", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	var list = utils.ReadInputFileToInts("input.txt")

	result := PartTwo(list, 25)
	expected := 55732936
	if result != expected {
		t.Errorf("Day 07, Part 2: %d; want %d", result, expected)
	}
}
