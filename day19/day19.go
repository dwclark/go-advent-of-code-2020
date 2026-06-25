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
		return false, -2
	} else {
		return toMatch[at] == rune(e), at + 1
	}
}

type either struct {
	couldBe [][]int64
}

func (e either) matches(rules map[int64]rule, toMatch []rune, at int) (bool, int) {
	if len(toMatch) <= at {
		return false, -2
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

type recurser func(level int) []int64

func drive(allowRecurse bool, r1 recurser, r2 recurser, rules map[int64]rule, toMatch []rune) bool {
	i := 0
r1Loop:
	for {
		r1matched, r1index := false, 0
		i++
		r1next := r1(i)
		if len(toMatch) <= len(r1next) {
			return false
		}

		for _, num := range r1next {
			rule := rules[num]
			r1matched, r1index = rule.matches(rules, toMatch, r1index)
			if !r1matched {
				if r1index == -2 {
					return false
				} else {
					continue r1Loop
				}
			}
		}

		j := 0
	r2Loop:
		for {
			r2matched, r2index := true, r1index
			j++
			r2next := r2(j)
			if len(toMatch)-r1index <= len(r2next) {
				break r2Loop
			}

			for _, num := range r2(j) {
				rule := rules[num]
				r2matched, r2index = rule.matches(rules, toMatch, r2index)
				if !r2matched {
					if r2index == -2 {
						break r2Loop
					} else {
						continue r2Loop
					}
				}
			}

			if r2matched && r2index == len(toMatch) {
				return true
			}
		}
	}
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

func Part2() int64 {
	rules, runes := parse("inputs/day-19.txt")

	f := func(toDup []int64) recurser {
		return func(times int) []int64 {
			ret := []int64{}
			for _, val := range toDup {
				for i := 0; i < times; i++ {
					ret = append(ret, val)
				}
			}

			return ret
		}
	}

	matched := int64(0)
	for _, ary := range runes {
		matches := drive(true, f([]int64{42}), f([]int64{42, 31}), rules, ary)
		if matches {
			matched++
		}
	}

	return utils.TestResult(19, 2, 369, matched)
}
