package day02

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

func GetNums() []int64 {
	return array.Map(utils.ToInt64)(utils.ReadLines("inputs/day-02.txt"))
}

func Part1() int64 {
	return utils.TestResult(2, 1, 956091, 0)
}

func Part2() int64 {
	return utils.TestResult(2, 2, 79734368, 0)
}
