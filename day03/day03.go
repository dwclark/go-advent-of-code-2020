package day03

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

func slopes() [][]rune {
	return array.Map(func(s string) []rune { return []rune(s) })(utils.ReadLines("inputs/day-03.txt"))
}

func encounters(all [][]rune, rightBy, downBy int) int64 {
	var result int64
	size := len(all[0])
	at := 0
	for row := 0; row < len(all); row += downBy {
		if all[row][at%size] == '#' {
			result++
		}

		at += rightBy
	}

	return result
}

func Part1() int64 {
	return utils.TestResult(3, 1, 145, encounters(slopes(), 3, 1))
}

func Part2() int64 {
	all := slopes()
	result := encounters(all, 1, 1) * encounters(all, 3, 1) *
		encounters(all, 5, 1) * encounters(all, 7, 1) * encounters(all, 1, 2)

	return utils.TestResult(3, 2, 3424528800, result)
}
