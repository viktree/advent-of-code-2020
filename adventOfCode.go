package main

import (
	"fmt"

	dayOne "github.com/viktree/advent-of-code-2020/day01"
	"github.com/viktree/advent-of-code-2020/utils"
)

func main() {
	var list = utils.ReadInputFileToInts("day01/input.txt")
	n := dayOne.PartOne(list)
	fmt.Printf("Answer %d\n", n)
}
