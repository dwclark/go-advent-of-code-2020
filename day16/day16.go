package day16

import (
	"aoc-2020/utils"
	//"fmt"
	"github.com/IBM/fp-go/array"
	"regexp"
	"strings"
)

type interval struct {
	low, high int64
}

func (i interval) inRange(val int64) bool {
	return i.low <= val && val <= i.high
}

func parse() (map[string][]interval, []int64, [][]int64) {
	fields := map[string][]interval{}
	var your []int64
	nearby := [][]int64{}

	mode := 1
	findRanges := regexp.MustCompile(`([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	lines := utils.ReadLines("inputs/day-16.txt")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			mode++
			i++
			continue
		}

		switch mode {
		case 1: //ranges
			matches := findRanges.FindStringSubmatch(line)
			fields[matches[1]] = []interval{interval{utils.ToInt64(matches[2]), utils.ToInt64(matches[3])},
				interval{utils.ToInt64(matches[4]), utils.ToInt64(matches[5])}}
		case 2: //your ticket
			your = array.Map(utils.ToInt64)(strings.Split(line, ","))
		case 3: //nearby tickets
			nearby = append(nearby, array.Map(utils.ToInt64)(strings.Split(line, ",")))
		}
	}

	return fields, your, nearby
}

func validInSomeRange(fields map[string][]interval, val int64) bool {
	for _, intervals := range fields {
		for _, interval := range intervals {
			if interval.inRange(val) {
				return true
			}
		}
	}

	return false
}

func invalidSum(fields map[string][]interval, nearby [][]int64) int64 {
	var ret int64
	for _, testArray := range nearby {
		for _, toTest := range testArray {
			if !validInSomeRange(fields, toTest) {
				ret += toTest
			}
		}
	}

	return ret
}

func Part1() int64 {
	fields, _, nearby := parse()
	return utils.TestResult(16, 1, 23036, invalidSum(fields, nearby))
}

func Part2() int64 {
	return utils.TestResult(16, 2, 0, 0)
}
