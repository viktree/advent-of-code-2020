package dayEight

func tweakInstruction(instruction string) string {
	instructionType := instruction[0:3]
	if instructionType == "nop" {
		return "jmp " + instruction[4:]
	} else if instructionType == "jmp" {
		return "nop " + instruction[4:]
	}
	return instruction
}

func PartTwo(instructions []string) int {
	nInstructions := len(instructions)
	instructionSet := make(map[int]string)

	pc := 0
	acc := 0

	var step string

	for pc < nInstructions {
		if _, ok := instructionSet[pc]; ok {
			break
		} else {
			instructionSet[pc] = instructions[pc]
		}
		pc, acc = exec(instructions[pc], pc, acc)
	}

	for tweakIdx, originalInstruction := range instructionSet {
		if originalInstruction[0:3] == "acc" {
			continue
		}
		pc = 0
		acc = 0
		cache := make(map[int]bool)
		for pc < nInstructions {
			if pc == tweakIdx {
				step = tweakInstruction(originalInstruction)
			} else {
				step = instructions[pc]
			}
			if _, ok := cache[pc]; ok {
				break
			} else {
				cache[pc] = true
			}
			pc, acc = exec(step, pc, acc)
		}
		if pc == nInstructions {
			return acc
		}
	}
	return acc
}
