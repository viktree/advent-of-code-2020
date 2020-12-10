package dayNine

func PartTwo(numbers []int, amt int) int {
	listSize := len(numbers)

	var complement int
	var foundComplement bool
	var target int

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
			target = n
			break
		}
	}

	var j, sum, smallest, largest int
	for i := range numbers {
		j = i
		sum = 0
		smallest = target
		largest = -1
		for j < listSize && sum < target {
			s := numbers[j]
			sum += s
			if s < smallest {
				smallest = s
			} else if s > largest {
				largest = s
			}
			j++
		}
		if sum == target {
			return largest + smallest
		}
	}

	return listSize
}
