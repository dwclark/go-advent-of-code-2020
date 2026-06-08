package day10

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
	"slices"
)

func nums() []int64 {
	tmp := append(array.Map(utils.ToInt64)(utils.ReadLines("inputs/day-10.txt")), 0)
	var theMax int64
	for _, num := range tmp {
		theMax = max(theMax, num)
	}

	tmp = append(tmp, theMax+3)
	slices.Sort(tmp)
	return tmp
}

var all []int64 = nums()
var cache map[int]int64 = map[int]int64{}

func Part1() int64 {
	diffs := map[int64]int64{}
	for i := 1; i < len(all); i++ {
		diffs[all[i]-all[i-1]]++
	}

	return utils.TestResult(10, 1, 2760, diffs[1]*diffs[3])
}

func possibilities(index int) int64 {
	cachedVal, found := cache[index]
	if found {
		return cachedVal
	}

	val := all[index]
	var mine int64
	if index == len(all)-1 {
		mine = 1
	} else {
		for i := index + 1; i < len(all); i++ {
			nextVal := all[i]
			if nextVal-val <= 3 {
				mine += possibilities(i)
			} else {
				break
			}
		}
	}

	cache[index] = mine
	return mine
}

func Part2() int64 {
	clear(cache)
	return utils.TestResult(10, 2, 13816758796288, possibilities(0))
}
