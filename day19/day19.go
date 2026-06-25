package day19

import (
	"aoc-2020/utils"
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
	if len(toMatch) <= at {
		return false, -1
	}

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
	rules, runes := parse("inputs/day-19.txt")
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

func recurse(times int, pattern []int64) either {
	tmp := []int64{}
	for _, val := range pattern {
		for i := 0; i < times; i++ {
			tmp = append(tmp, val)
		}
	}

	return either{[][]int64{tmp}}
}

func execP2(p1 []int64, p2 []int64, rules map[int64]rule, toMatch []rune) bool {
	for frecurse := 1; frecurse*len(p1) <= len(toMatch); frecurse++ {
		first := recurse(frecurse, p1)
		fmatched, findex := first.matches(rules, toMatch, 0)
		if fmatched {
			for srecurse := 1; srecurse*len(p2) <= len(toMatch)-findex; srecurse++ {
				second := recurse(srecurse, p2)
				smatched, sindex := second.matches(rules, toMatch, findex)
				if smatched && sindex == len(toMatch) {
					return true
				}
			}
		}
	}

	return false
}

func Part2() int64 {
	rules, runes := parse("inputs/day-19.txt")
	matched := int64(0)
	for _, ary := range runes {
		matches := execP2([]int64{42}, []int64{42, 31}, rules, ary)
		if matches {
			matched++
		}
	}

	return utils.TestResult(19, 2, 369, matched)
}
