package day05

import (
	"aoc-2020/utils"
)

func seats() []string {
	return utils.ReadLines("inputs/day-05.txt")
}

func find(line string, lowRune rune, lower int, upRune rune, upper int) int64 {
	for _, r := range line {
		sum := upper + lower
		div := sum / 2
		mod := sum % 2
		if r == lowRune {
			upper = div
		} else if r == upRune {
			lower = div + mod
		}
	}

	return int64(lower)
}

func row(line string) int64 {
	return find(line, 'F', 0, 'B', 127)
}

func col(line string) int64 {
	return find(line, 'L', 0, 'R', 7)
}

func Part1() int64 {
	var result int64
	for _, line := range seats() {
		result = max(((row(line) * 8) + col(line)), result)
	}

	return utils.TestResult(5, 1, 816, result)
}

func Part2() int64 {
	all := [1024]bool{}
	var result int64

	for _, line := range seats() {
		all[((row(line) * 8) + col(line))] = true
	}

	for i := 1; i < 1024; i++ {
		if all[i-1] && !all[i] && all[i+1] {
			result = int64(i)
			break
		}
	}

	return utils.TestResult(5, 2, 539, result)
}
