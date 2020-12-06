package dayFive

import (
	"fmt"
	"strconv"
)

func uintToInt(u uint64) int {
	s := fmt.Sprint(u)
	r, _ := strconv.Atoi(s)
	return r
}

func PartTwo(boardingpasses []string) int {
	n := len(boardingpasses)
	whatSumIs := 0
	largest := 0
	for _, boardingpass := range boardingpasses {
		row := getRow(boardingpass)
		col := getCol(boardingpass)
		num := uintToInt(row<<3 + col)
		if num > largest {
			largest = num
		}
		whatSumIs += num
	}
	whatSumShouldBe := largest*(n+1) - n*(n+1)/2
	return whatSumShouldBe - whatSumIs
}
