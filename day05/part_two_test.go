package dayFive

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPartTwo(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartTwo(list)
	expected := 519
	if result != expected {
		t.Errorf("Day 5, Part 2: %d; want %d", result, expected)
	}
}
