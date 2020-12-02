package dayOne

func PartOne(list []int) int {
	set := make(map[int]bool)

	for _, n := range list {
		if _, ok := set[2020-n]; ok {
			return n * (2020 - n)
		}
		set[n] = true
	}
	return -1
}
