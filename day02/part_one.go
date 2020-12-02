package dayTwo

import (
	"fmt"
	"strings"
)

func PartOne(list []string) int {

	var minAmount, maxAmount, count int
	var letter, password string
	nVaild := 0
	for _, rawString := range list {

		var splits = strings.Split(rawString, ":")
		password = splits[1]
		_, err := fmt.Sscanf(splits[0], "%d-%d %s", &minAmount, &maxAmount, &letter)
		if err != nil {
			fmt.Printf("Failed to scan %s\n", rawString)
		}

		count = 0
		for _, l := range password {
			if string(l) == letter {
				count++
			}
		}
		if count >= minAmount && count <= maxAmount {
			nVaild++
		}
	}
	return nVaild
}
