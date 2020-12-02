package dayOne

import (
	"fmt"
	"strconv"

	"github.com/viktree/advent-of-code-2020/utils"
)

func PartOne() {

	var input = utils.ReadInputFile("01", "input.txt")
	set := make(map[int]bool)

	var numAsString string
	for _, numAsString = range input {
		n, _ := strconv.Atoi(numAsString)
		if _, ok := set[2020-n]; ok {
			fmt.Printf("Answer %d\n", (n)*(2020-n))
			return
		}
		set[n] = true
	}
}

func PartTwo() {
	var input = utils.ReadInputFile("01", "input.txt")
	set := make(map[int]bool)

	var xAsString, yAsString string
	var x, y int

	for _, xAsString = range input {
		x, _ = strconv.Atoi(xAsString)
		set[x] = true
	}

	for _, xAsString = range input {
		x, _ = strconv.Atoi(xAsString)
		for _, yAsString = range input {
			y, _ = strconv.Atoi(yAsString)
			if _, ok := set[2020-(x+y)]; ok {
				fmt.Printf("Answer %d\n", x*y*(2020-x-y))
				return
			}
		}
	}
}
