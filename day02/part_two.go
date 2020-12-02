package dayTwo

import (
	"fmt"
	"strings"
)

func PartTwo(list []string) int {

	var posOne, posTwo int
	var letter, password string
	nVaild := 0
	for _, rawString := range list {

		var splits = strings.Split(rawString, ": ")
		password = splits[1]
		_, err := fmt.Sscanf(splits[0], "%d-%d %s", &posOne, &posTwo, &letter)
		if err != nil {
			fmt.Printf("Failed to scan %s\n", rawString)
		}

		if posTwo > len(password) {
			continue
		}

		if (string(password[posOne-1]) == letter) != (string(password[posTwo-1]) == letter) {
			nVaild++
		}
	}
	return nVaild
}
