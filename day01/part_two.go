package dayOne

func PartTwo(list []int) int {
	set := make(map[int]bool)

	var x, y int

	for _, x = range list {
		set[x] = true
		for _, y = range list {
			if _, ok := set[2020-(x+y)]; ok {
				return x * y * (2020 - x - y)
			}
		}
	}

	return -1
}
