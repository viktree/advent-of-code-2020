package dayOne

import (
	"fmt"

	"github.com/viktree/advent-of-code-2020/utils"
)

func PartOne() {

	var list = utils.ReadInputFileToInts("01", "input.txt")
	set := make(map[int]bool)

	for _, n := range list {
		if _, ok := set[2020-n]; ok {
			fmt.Printf("Answer %d\n", (n)*(2020-n))
			return
		}
		set[n] = true
	}
}
