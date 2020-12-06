package dayFive

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartOneWithExample(t *testing.T) {
	var list = utils.ReadInputFile("example.txt")

	result := PartOne(list)
	expected := uint64(820)
	if result != expected {
		t.Errorf("Day 5, Part 1 (example): %d; want %d", result, expected)
	}
}

func TestPartOne(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartOne(list)
	expected := uint64(906)
	if result != expected {
		t.Errorf("Day 5, Part 1: %d; want %d", result, expected)
	}
}
