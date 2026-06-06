package day02

import (
	"aoc-2020/utils"
	"strings"
)

type datum struct {
	min      int64
	max      int64
	letter   rune
	password string
}

func data() []datum {
	var data []datum
	for _, line := range utils.ReadLines("inputs/day-02.txt") {
		all := strings.Split(line, " ")
		minMax := strings.Split(all[0], "-")
		data = append(data, datum{
			min:      utils.ToInt64(minMax[0]),
			max:      utils.ToInt64(minMax[1]),
			letter:   []rune(all[1])[0],
			password: all[2],
		})
	}

	return data
}

func Part1() int64 {
	var valid int64
	for _, d := range data() {
		var times int64
		for _, c := range d.password {
			if c == d.letter {
				times++
			}
		}

		if d.min <= times && times <= d.max {
			valid++
		}
	}

	return utils.TestResult(2, 1, 469, valid)
}

func Part2() int64 {
	var valid int64
	for _, d := range data() {
		pwd := []rune(d.password)
		size := int64(len(pwd))
		if d.max <= size && utils.Xor(pwd[d.min-1] == d.letter, pwd[d.max-1] == d.letter) {
			valid++
		}
	}

	return utils.TestResult(2, 2, 267, valid)
}
