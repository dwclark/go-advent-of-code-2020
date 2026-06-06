package day01

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
	"github.com/etnz/permute"
)

func GetNums() []int64 {
	return array.Map(utils.ToInt64)(utils.ReadLines("inputs/day-01.txt"))
}

func Part1() int64 {
	nums := GetNums()
	var result int64

outer:
	for idx, outer := range nums {
		for _, inner := range nums[idx+1:] {
			if outer+inner == 2020 {
				result = outer * inner
				break outer
			}
		}
	}

	return utils.TestResult(1, 1, 956091, result)
}

func Part2() int64 {
	var result int64
	for comb := range permute.Combinations(3, GetNums()) {
		if comb[0]+comb[1]+comb[2] == 2020 {
			result = comb[0] * comb[1] * comb[2]
			break
		}
	}

	return utils.TestResult(1, 2, 79734368, result)
}
