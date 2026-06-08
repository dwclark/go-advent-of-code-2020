package day09

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
	"github.com/etnz/permute"
	"slices"
)

func nums() []int64 {
	return array.Map(utils.ToInt64)(utils.ReadLines("inputs/day-09.txt"))
}

func sumPrevious(num int64, prev []int64) bool {
	for comb := range permute.Combinations(2, prev) {
		if num == (comb[0] + comb[1]) {
			return true
		}
	}

	return false
}

func Part1() int64 {
	var result int64
	all := nums()

	for i := 25; i < len(all); i++ {
		if !sumPrevious(all[i], all[(i-25):i]) {
			result = all[i]
			break
		}
	}

	return utils.TestResult(9, 1, 177777905, result)
}

func findRange(sumTo int64, all []int64) (int, int) {
	for outer := 0; outer < len(all); outer++ {
		var sum int64 = all[outer]
		for inner := outer + 1; inner < len(all); inner++ {
			sum += all[inner]
			if sum < sumTo {
				continue
			} else if sum == sumTo {
				return outer, inner + 1
			} else {
				break
			}
		}
	}

	return -1, -1
}

func Part2() int64 {
	all := nums()
	sumTo := Part1()
	lower, upper := findRange(sumTo, all)
	slice := all[lower:upper]
	slices.Sort(slice)
	return utils.TestResult(9, 2, 23463012, slice[0]+slice[len(slice)-1])
}
