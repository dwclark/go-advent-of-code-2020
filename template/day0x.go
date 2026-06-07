package day0x

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

func nums() []int64 {
	return array.Map(utils.ToInt64)(utils.ReadLines("inputs/day-0x.txt"))
}

func Part1() int64 {
	return utils.TestResult(-1, 1, 0, 0)
}

func Part2() int64 {
	return utils.TestResult(-1, 2, 0, 0)
}
