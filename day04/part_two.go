package dayFour

import (
	"regexp"
	"strconv"
	"strings"
)

func validatePassportKey(key string, value string) bool {
	var matched bool
	switch key {
	case "cid":
		matched = true
	case "byr":
		year, _ := strconv.Atoi(value)
		matched = year >= 1920 && year <= 2002
	case "iyr":
		year, _ := strconv.Atoi(value)
		matched = year >= 2010 && year <= 2020
	case "eyr":
		year, _ := strconv.Atoi(value)
		matched = year >= 2020 && year <= 2030
	case "pid":
		matched, _ = regexp.MatchString(`^\d\d\d\d\d\d\d\d\d$`, value)
	case "hcl":
		matched, _ = regexp.MatchString(`^#([a-f]|\d)([a-f]|\d)([a-f]|\d)([a-f]|\d)([a-f]|\d)([a-f]|\d)$`, value)
	case "ecl":
		matched, _ = regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, value)
	case "hgt":
		matchedCm, _ := regexp.MatchString(`^(\d\d\d)(cm)$`, value)
		matchedIn, _ := regexp.MatchString(`^(\d\d)(in)$`, value)
		if matchedCm {
			cm, _ := strconv.Atoi(value[0:3])
			matched = cm >= 150 && cm <= 193
		} else if matchedIn {
			in, _ := strconv.Atoi(value[0:2])
			matched = in >= 59 && in <= 76
		} else {
			matched = false
		}
	default:
		matched = false
	}
	return matched
}

func PartTwo(list []string) int {
	var matched, isValid bool

	count := 0
	passports := getPassports(list)

	for _, passport := range passports {
		if !passportIsValid(passport) {
			continue
		}
		entries := strings.Split(passport, " ")
		isValid = true
		for _, entry := range entries {
			if entry == "" {
				continue
			}
			keyVal := strings.Split(entry, ":")
			matched = validatePassportKey(keyVal[0], keyVal[1])

			if !matched {
				isValid = false
				break
			}
		}
		if isValid {
			count++
		}
	}

	return count
}
