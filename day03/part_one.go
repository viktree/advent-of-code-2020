package dayThree

func PartOne(list []string) int {
	nRows := len(list[0])

	count := 0
	for row_idx, row := range list {
		x := 3 * row_idx % nRows
		for j := range row {
			if j == x && row[x] == '#' {
				count++
			}
		}
	}
	return count
}
