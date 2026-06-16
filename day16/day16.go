package day16

import (
	"aoc-2020/utils"
	//"fmt"
	"github.com/IBM/fp-go/array"
	"maps"
	"regexp"
	"strings"
)

type intervals struct {
	low1, high1, low2, high2 int64
}

func (i intervals) inRange(val int64) bool {
	return (i.low1 <= val && val <= i.high1) || (i.low2 <= val && val <= i.high2)
}

func parse() (map[string]intervals, []int64, [][]int64) {
	fields := map[string]intervals{}
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
			fields[matches[1]] = intervals{utils.ToInt64(matches[2]), utils.ToInt64(matches[3]),
				utils.ToInt64(matches[4]), utils.ToInt64(matches[5])}
		case 2: //your ticket
			your = array.Map(utils.ToInt64)(strings.Split(line, ","))
		case 3: //nearby tickets
			nearby = append(nearby, array.Map(utils.ToInt64)(strings.Split(line, ",")))
		}
	}

	return fields, your, nearby
}

func validInSomeRange(fields map[string]intervals, val int64) bool {
	for _, toTest := range fields {
		if toTest.inRange(val) {
			return true
		}
	}

	return false
}

func invalidSum(fields map[string]intervals, nearby [][]int64) int64 {
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

func valids(fields map[string]intervals, nearby [][]int64) [][]int64 {
	ret := [][]int64{}
outerLoop:
	for _, testArray := range nearby {
		for _, toTest := range testArray {
			if !validInSomeRange(fields, toTest) {
				continue outerLoop
			}
		}

		ret = append(ret, testArray)
	}

	return ret
}

func validPosition(toTest intervals, position int, tickets [][]int64) bool {
	for _, ticket := range tickets {
		if !toTest.inRange(ticket[position]) {
			return false
		}
	}

	return true
}

func satisfy(fields map[string]intervals, positions []string, tickets [][]int64) (bool, []string) {
	if len(positions) == len(tickets[0]) {
		return true, positions
	}

	for field, toTest := range fields {
		if validPosition(toTest, len(positions), tickets) {
			newFields := maps.Clone(fields)
			delete(newFields, field)
			newPositions := append(positions, field)
			isValid, newPositions := satisfy(newFields, newPositions, tickets)
			if isValid {
				return isValid, newPositions
			}
		}
	}

	return false, []string{} //keep compiler happy, should not get here
}

func Part1() int64 {
	fields, _, nearby := parse()
	return utils.TestResult(16, 1, 23036, invalidSum(fields, nearby))
}

func Part2() int64 {
	fields, yours, nearby := parse()
	allValid := valids(fields, nearby)
	isValid, positions := satisfy(fields, []string{}, allValid)

	var result int64 = int64(1)
	if isValid {
		for idx, name := range positions {
			if strings.HasPrefix(name, "departure") {
				result *= yours[idx]
			}
		}
	}

	return utils.TestResult(16, 2, 1909224687553, result)
}
