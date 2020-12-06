package dayFive

import (
	"strconv"
	"strings"
)

func getRow(boardingpass string) uint64 {
	releventPart := boardingpass[0:7]
	releventPart = strings.Replace(releventPart, "B", "1", 8)
	releventPart = strings.Replace(releventPart, "F", "0", 8)
	u, err := strconv.ParseUint(releventPart, 2, 32)
	if err != nil {
		panic(err)
	}
	return u
}

func getCol(boardingpass string) uint64 {
	releventPart := boardingpass[7:]
	releventPart = strings.Replace(releventPart, "R", "1", 3)
	releventPart = strings.Replace(releventPart, "L", "0", 3)
	u, err := strconv.ParseUint(releventPart, 2, 4)
	if err != nil {
		panic(err)
	}
	return u
}

func getSeatId(boardingpass string) uint64 {
	row := getRow(boardingpass)
	col := getCol(boardingpass)
	return row*8 + col

}

func PartOne(boardingpasses []string) uint64 {
	var seatId, highestSeatIdSoFar uint64
	highestSeatIdSoFar = uint64(0)
	for _, boardingpass := range boardingpasses {
		seatId = getSeatId(boardingpass)
		if seatId > highestSeatIdSoFar {
			highestSeatIdSoFar = seatId
		}
	}

	return highestSeatIdSoFar
}
