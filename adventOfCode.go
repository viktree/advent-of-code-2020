package main

import (
	"fmt"

	dayFour "github.com/viktree/advent-of-code-2020/day04"
	"github.com/viktree/advent-of-code-2020/utils"
)

func main() {
	var list = utils.ReadInputFile("day04/input.txt")
	n := dayFour.PartTwo(list)
	fmt.Printf("Answer %d\n", n)
}
