package day07

import (
	"aoc-2020/utils"
	"strings"
	//"github.com/IBM/fp-go/array"
	//"fmt"
)

type bag struct {
	color    string
	contains map[string]int64
}

func parseBag(line string) bag {
	allParts := strings.Split(line, " ")
	ret := bag{color: allParts[0] + " " + allParts[1], contains: map[string]int64{}}
	for i := 4; i < len(allParts); i += 4 {
		if allParts[i] == "no" {
			return ret
		}

		count := utils.ToInt64(allParts[i])
		color := allParts[i+1] + " " + allParts[i+2]
		ret.contains[color] = count
	}

	return ret
}

func bags() map[string]bag {
	ret := map[string]bag{}
	for _, line := range utils.ReadLines("inputs/day-07.txt") {
		theBag := parseBag(line)
		ret[theBag.color] = theBag
	}

	return ret
}

func containsShinyGold(theBag bag, bags map[string]bag) bool {
	if theBag.contains["shiny gold"] > 0 {
		return true
	} else if len(theBag.contains) == 0 {
		return false
	} else {
		for newColor, _ := range theBag.contains {
			if containsShinyGold(bags[newColor], bags) {
				return true
			}
		}

		return false
	}
}

func Part1() int64 {
	var result int64
	all := bags()
	for _, theBag := range all {
		if containsShinyGold(theBag, all) {
			result++
		}
	}

	return utils.TestResult(7, 1, 246, result)
}

func Part2() int64 {
	return utils.TestResult(7, 2, 0, 0)
}
