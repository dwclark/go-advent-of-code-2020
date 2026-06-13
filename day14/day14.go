package day14

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
	"github.com/etnz/permute"
	"slices"
	"strings"
)

type mem struct {
	address, val uint64
}

func toMask(s string) []rune {
	tmp := []rune(s)
	for left, right := 0, len(tmp)-1; left < right; left, right = left+1, right-1 {
		tmp[left], tmp[right] = tmp[right], tmp[left]
	}

	return tmp
}

func parse(s string) any {
	if strings.HasPrefix(s, "mask") {
		return toMask(strings.Trim(strings.Split(s, "=")[1], " "))
	} else {
		parts := strings.Split(s, "=")
		memParts := strings.Split(strings.ReplaceAll(parts[0], "]", ""), "[")
		return mem{utils.ToUInt64(strings.Trim(memParts[1], " ")),
			utils.ToUInt64(strings.Trim(parts[1], " "))}
	}
}

func instructions() []any {
	return array.Map(parse)(utils.ReadLines("inputs/day-14.txt"))
}

type handleMem func(i mem, mask []rune, ram map[uint64]uint64)

func run(f handleMem) int64 {
	var mask []rune
	ram := map[uint64]uint64{}

	for _, instruction := range instructions() {
		switch t := instruction.(type) {
		case []rune:
			mask = t
		case mem:
			f(t, mask, ram)
		}
	}

	var result uint64
	for _, v := range ram {
		result += v
	}

	return int64(result)
}

func floating(runes []rune) []int {
	indexes := []int{}

	for i := 0; i < len(runes); i++ {
		if runes[i] == 'X' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func allPossible(indexes []int) [][]int {
	powerSet := [][]int{}
	for comboIndex := 0; comboIndex <= len(indexes); comboIndex++ {
		for combo := range permute.Combinations(comboIndex, indexes) {
			powerSet = append(powerSet, combo)
		}
	}

	return powerSet
}

func Part1() int64 {
	f := func(i mem, mask []rune, ram map[uint64]uint64) {
		var accum uint64
		for idx, r := range mask {
			if r == '1' {
				accum |= (1 << idx)
			} else if r == '0' {
				accum = accum & ^(1 << idx)
			} else if r == 'X' {
				accum |= (i.val & (1 << idx))
			}
		}

		ram[i.address] = accum
	}

	return utils.TestResult(14, 1, 14722016054794, run(f))
}

func Part2() int64 {
	f := func(i mem, mask []rune, ram map[uint64]uint64) {
		//apply _all_ non-floating
		masked := i.address
		for idx, r := range mask {
			if r == '1' {
				masked |= (1 << idx)
			}
		}

		//now apply floating
		allFloating := floating(mask)
		for _, set := range allPossible(allFloating) {
			accum := masked
			for _, idx := range allFloating {
				if slices.Contains(set, idx) {
					accum |= (1 << idx)
				} else {
					accum = accum & ^(1 << idx)
				}
			}

			ram[accum] = i.val
		}
	}

	return utils.TestResult(14, 2, 3618217244644, run(f))
}
