package dayThree

func PartOne(list []string) int {
	nRows := len(list[0])

	count := 0
	for row_idx, row := range list {
		if row[3*row_idx%nRows] == '#' {
			count++
		}
	}
	return count
}
