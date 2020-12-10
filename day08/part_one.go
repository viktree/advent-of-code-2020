package dayEight

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(instructions []string) int {

	var rawInstruction []string
	seenInstructionBefore := make(map[int]bool)
	var instruction string
	i := 0
	acc := 0
	for i < len(instructions) {
		if _, ok := seenInstructionBefore[i]; ok {
			break
		} else {
			seenInstructionBefore[i] = true
		}
		rawInstruction = strings.Split(instructions[i], " ")
		instruction = rawInstruction[0]
		value, _ := strconv.Atoi(rawInstruction[1])
		switch instruction {
		case "jmp":
			fmt.Println("JUMP")
			i += value
		case "acc":
			fmt.Println("ACC")
			acc += value
			i++
		case "nop":
			fmt.Println("NOP")
			i++
		}
	}
	return acc
}
