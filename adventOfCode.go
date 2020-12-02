package main

import (
	"fmt"

	dayTwo "github.com/viktree/advent-of-code-2020/day02"
	"github.com/viktree/advent-of-code-2020/utils"
)

func main() {
	var list = utils.ReadInputFile("day02/input.txt")
	n := dayTwo.PartOne(list)
	fmt.Printf("Answer %d\n", n)
}
