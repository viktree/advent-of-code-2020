package dayFour

import (
	"strings"
)

func getPassports(list []string) []string {
	var passports []string
	start := 0
	for start < len(list) {
		passport := ""
		for _, character := range list[start:] {
			start++
			if string(character) == string("") {
				passport += " "
				break
			}
			passport += character + " "
		}
		passports = append(passports, passport)
	}
	return passports
}

func passportIsValid(passport string) bool {
	hasByr := strings.Contains(passport, "byr")
	hasIyr := strings.Contains(passport, "iyr")
	hasEyr := strings.Contains(passport, "eyr")
	hasHgt := strings.Contains(passport, "hgt")
	hasHcl := strings.Contains(passport, "hcl")
	hasEcl := strings.Contains(passport, "ecl")
	hasPid := strings.Contains(passport, "pid")
	hasKeys := hasIyr && hasPid && hasEyr && hasHcl && hasHgt && hasByr && hasEcl
	return hasKeys
}

func PartOne(list []string) int {
	count := 0
	passports := getPassports(list)
	for _, potentialPassport := range passports {
		if passportIsValid(potentialPassport) {
			count++
		}
	}

	return count
}
