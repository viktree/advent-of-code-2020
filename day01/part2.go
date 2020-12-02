package dayOne

import (
	"fmt"

	"github.com/viktree/advent-of-code-2020/utils"
)

func PartTwo() {
	var list = utils.ReadInputFileToInts("01", "input.txt")
	set := make(map[int]bool)

	var x, y int

	for _, x = range list {
		set[x] = true
		for _, y = range list {
			if _, ok := set[2020-(x+y)]; ok {
				fmt.Printf("Answer %d\n", x*y*(2020-x-y))
				return
			}
		}
	}
}
