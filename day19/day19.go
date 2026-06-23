package day19

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

type rule interface {
	matches(rules map[int64]rule, toMatch []rune, at int) (bool, int)
}

type exact rune

func (e exact) matches(rules map[int64]rule, toMatch []rune, at int) (bool, int) {
	if len(toMatch) <= at {
		return false, -1
	} else {
		return toMatch[at] == rune(e), at + 1
	}
}

type either struct {
	couldBe [][]int64
}

func (e either) matches(rules map[int64]rule, toMatch []rune, at int) (bool, int) {
	for _, nums := range e.couldBe {
		matched, index := false, at
		for _, num := range nums {
			rule := rules[num]
			matched, index = rule.matches(rules, toMatch, index)
			if !matched {
				break
			}
		}

		if matched {
			return matched, index
		}
	}

	return false, -1
}

func parseRule(line string) (int64, rule) {
	ary := strings.Split(line, ":")
	num := utils.ToInt64(ary[0])
	if strings.Contains(ary[1], "\"") {
		return num, exact(strings.Trim(strings.ReplaceAll(ary[1], "\"", ""), " ")[0])
	} else {
		strParts := strings.Split(strings.Trim(ary[1], " "), "|")
		couldBe := [][]int64{}
		for _, sub := range strParts {
			tmp := []int64{}
			for _, strNum := range strings.Split(strings.Trim(sub, " "), " ") {
				tmp = append(tmp, utils.ToInt64(strings.Trim(strNum, " ")))
			}

			couldBe = append(couldBe, tmp)
		}

		return num, either{couldBe}
	}
}

func parse(file string) (map[int64]rule, [][]rune) {
	rules := map[int64]rule{}
	runes := [][]rune{}
	mode := 1
	for _, line := range utils.ReadLines(file) {
		if len(line) == 0 {
			mode++
			continue
		}

		if mode == 1 {
			num, rule := parseRule(line)
			rules[num] = rule
		} else {
			runes = append(runes, []rune(line))
		}
	}

	return rules, runes
}

func Part1() int64 {
	rules, runes := parse("inputs/day-19-1.txt")
	matched := int64(0)
	zero := rules[int64(0)]

	for _, ary := range runes {
		matches, index := zero.matches(rules, ary, 0)
		if matches && index == len(ary) {
			matched++
		}
	}

	return utils.TestResult(19, 1, 171, matched)
}

func Part2() int64 {
	rules, runes := parse("inputs/day-19-b.txt")
	matched := int64(0)
	zero := rules[int64(0)]

	for _, ary := range runes {
		matches, index := zero.matches(rules, ary, 0)
		if matches && index == len(ary) {
			fmt.Println("Matched", string(ary))
			matched++
		}
	}

	return utils.TestResult(19, 2, 12, matched)
}
