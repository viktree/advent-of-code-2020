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
	product := make(map[int]int)

	var xAsString, yAsString, zAsString string
	var x, y, z int

	for _, xAsString = range input {
		x, _ = strconv.Atoi(xAsString)
		for _, yAsString = range input {
			y, _ = strconv.Atoi(yAsString)
			z = 2020 - (x + y)
			set[z] = true
			product[z] = x * y * z
		}
	}

	for _, zAsString = range input {
		z, _ = strconv.Atoi(zAsString)
		if _, ok := set[z]; ok {
			fmt.Printf("Answer %d\n", product[z])
			return
		}
	}
}
