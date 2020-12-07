package daySix

type Answer struct {
	response rune
	group    int
}

func PartOne(peopleResponses []string) int {
	ngroup := 0
	responeses := make(map[Answer]bool)
	for _, line := range peopleResponses {
		if line == "" {
			ngroup++
			continue
		}
		for _, answer := range line {
			responeses[Answer{response: answer, group: ngroup}] = true
		}
	}
	return len(responeses)
}

func PartTwo(peopleResponses []string) int {
	ngroup := 1
	personIdx := 0
	resCount := make(map[rune]int)
	count := 0
	for i, line := range peopleResponses {
		if line == "" {
			for _, v := range resCount {
				if v == personIdx {
					count++
				}
			}
			ngroup++
			personIdx = 0
			resCount = make(map[rune]int)
			continue
		}
		for _, answer := range line {
			if _, ok := resCount[answer]; ok {
				resCount[answer] += 1
			} else if i == 0 || peopleResponses[i-1] == "" {
				resCount[answer] = 1
			}
		}
		personIdx++
	}
	for _, v := range resCount {
		if v == personIdx {
			count++
		}
	}
	return count
}
