package dayTwo

import (
	"testing"

	"github.com/viktree/advent-of-code-2020/utils"
)

func TestPart1(t *testing.T) {
	var list = utils.ReadInputFile("input.txt")

	result := PartOne(list)
	expected := 560
	if result != expected {
		t.Errorf("PartOne(): %d; want %d", result, expected)
	}
}
