package dayEight

import (
	"strconv"
	"strings"
)

func exec(instruction string, programCounter int, accumulator int) (int, int) {
	parts := strings.Split(instruction, " ")
	value, _ := strconv.Atoi(parts[1])
	switch parts[0] {
	case "jmp":
		programCounter += value
	case "acc":
		accumulator += value
		programCounter++
	case "nop":
		programCounter++
	}
	return programCounter, accumulator
}

func PartOne(instructions []string) int {
	instructionSet := make(map[int]bool)
	nInstructions := len(instructions)

	pc := 0
	acc := 0
	for pc < nInstructions {
		if _, ok := instructionSet[pc]; ok {
			break
		} else {
			instructionSet[pc] = true
		}
		pc, acc = exec(instructions[pc], pc, acc)
	}
	return acc
}
