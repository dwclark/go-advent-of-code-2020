package day04

import (
	"aoc-2020/utils"
	"log"
	"regexp"
	"strings"
)

func passports() []map[string]string {
	ret := []map[string]string{}
	current := map[string]string{}

	for _, line := range utils.ReadLines("inputs/day-04.txt") {
		if line == "" {
			ret = append(ret, current)
			current = map[string]string{}
			continue
		}

		for _, entry := range strings.Split(line, " ") {
			keyVal := strings.Split(entry, ":")
			current[keyVal[0]] = keyVal[1]
		}
	}

	if len(current) != 0 {
		ret = append(ret, current)
	}

	return ret
}

func bulkValid(passport map[string]string) bool {
	return len(passport) == 8 || (len(passport) == 7 && passport["cid"] == "")
}

func byr(passport map[string]string) bool {
	year := utils.ToInt64(passport["byr"])
	return 1920 <= year && year <= 2002
}

func iyr(passport map[string]string) bool {
	year := utils.ToInt64(passport["iyr"])
	return 2010 <= year && year <= 2020
}

func eyr(passport map[string]string) bool {
	year := utils.ToInt64(passport["eyr"])
	return 2020 <= year && year <= 2030
}

func hgt(passport map[string]string) bool {
	original := passport["hgt"]
	if strings.Contains(original, "cm") {
		val := utils.ToInt64(strings.ReplaceAll(original, "cm", ""))
		return 150 <= val && val <= 193
	} else if strings.Contains(original, "in") {
		val := utils.ToInt64(strings.ReplaceAll(original, "in", ""))
		return 59 <= val && val <= 76
	} else {
		return false
	}
}

func hcl(passport map[string]string) bool {
	matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, passport["hcl"])
	if err != nil {
		log.Fatal(err)
	}

	return matched
}

func ecl(passport map[string]string) bool {
	matched, err := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, passport["ecl"])
	if err != nil {
		log.Fatal(err)
	}

	return matched
}

func pid(passport map[string]string) bool {
	matched, err := regexp.MatchString(`^[0-9]{9}$`, passport["pid"])
	if err != nil {
		log.Fatal(err)
	}

	return matched
}

func Part1() int64 {
	var result int64
	for _, passport := range passports() {
		if bulkValid(passport) {
			result++
		}
	}

	return utils.TestResult(4, 1, 235, result)
}

func Part2() int64 {
	var result int64
	for _, passport := range passports() {
		if bulkValid(passport) && byr(passport) && iyr(passport) && eyr(passport) &&
			hgt(passport) && hcl(passport) && ecl(passport) && pid(passport) {
			result++
		}
	}

	return utils.TestResult(4, 2, 194, result)
}
