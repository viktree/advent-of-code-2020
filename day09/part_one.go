package dayNine

func PartOne(numbers []int, amt int) int {
	idx := 0

	var complement int
	var foundComplement bool
	for i, n := range numbers {
		if i < amt {
			continue
		}
		list := numbers[i-amt : i]
		set := make(map[int]bool)
		foundComplement = false
		for _, x := range list {
			if _, ok := set[x]; ok {
				foundComplement = true
				break
			}
			complement = n - x
			set[complement] = true
		}
		if !foundComplement {
			return n
		}
	}

	return idx
}
