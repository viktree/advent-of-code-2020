package dayThree

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartOne(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartOne(list)
	expected := 173
	if result != expected {
		t.Errorf("Day 3, Part 1: %d; want %d", result, expected)
	}
}
