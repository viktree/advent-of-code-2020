package dayNine

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartOneWithExample(t *testing.T) {
	var list = utils.ReadInputFileToInts("example.txt")

	result := PartOne(list, 5)
	expected := 127
	if result != expected {
		t.Errorf("Day 8, Part 1 (example): %d; want %d", result, expected)
	}
}

func TestPartOne(t *testing.T) {
	var list = utils.ReadInputFileToInts("input.txt")

	result := PartOne(list, 25)
	expected := 466456641
	if result != expected {
		t.Errorf("Day 8, Part 1: %d; want %d", result, expected)
	}
}
