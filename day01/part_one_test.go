package dayOne

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPart1(t *testing.T) {
	var list = utils.ReadInputFileToInts("input.txt")

	result := PartOne(list)
	expected := 471019
	if result != expected {
		t.Errorf("PartOne(): %d; want %d", result, expected)
	}
}
