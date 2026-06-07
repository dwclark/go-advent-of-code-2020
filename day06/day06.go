package day06

import (
	"aoc-2020/utils"
)

func lines() []string {
	return utils.ReadLines("inputs/day-06.txt")
}

func Part1() int64 {
	var result int
	group := map[rune]bool{}

	for _, line := range lines() {
		if len(line) > 0 {
			for _, r := range line {
				group[r] = true
			}
		} else {
			result += len(group)
			clear(group)
		}
	}

	result += len(group)
	return utils.TestResult(5, 1, 7283, int64(result))
}

func allAnswered(group map[rune]int, groupLines int) int {
	result := 0
	for _, val := range group {
		if val == groupLines {
			result++
		}
	}

	return result
}

func Part2() int64 {
	var result int
	group := map[rune]int{}
	groupLines := 0

	for _, line := range lines() {
		if len(line) > 0 {
			groupLines++
			for _, r := range line {
				group[r]++
			}
		} else {
			result += allAnswered(group, groupLines)
			clear(group)
			groupLines = 0
		}
	}

	result += allAnswered(group, groupLines)
	return utils.TestResult(5, 2, 3520, int64(result))
}
