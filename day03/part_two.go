package dayThree

func PartTwo(list []string) int {
	nRows := len(list[0])

	answer := 1
	counters := [5]int{0, 0, 0, 0, 0}
	leftAmt := [5]int{1, 3, 5, 7, 1}
	downAmt := [5]int{1, 1, 1, 1, 2}

	for step := 0; step < 5; step++ {
		for row_idx, row := range list {

			// move down by skipping rows
			if row_idx%downAmt[step] != 0 {
				continue
			}

			// move accross by computing the coords
			x := ((row_idx * leftAmt[step]) / downAmt[step]) % nRows
			for j := range row {
				if j == x {
					if row[x] == '#' {
						counters[step]++
					}
				}
			}
		}
		answer *= counters[step]
	}
	return answer
}
