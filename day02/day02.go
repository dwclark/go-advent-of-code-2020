package day02

import (
	"aoc-2020/utils"
	"strings"
)

type Data struct {
	min      int64
	max      int64
	letter   rune
	password string
}

func GetData() []Data {
	var data []Data
	for _, line := range utils.ReadLines("inputs/day-02.txt") {
		all := strings.Split(line, " ")
		minMax := strings.Split(all[0], "-")
		data = append(data, Data{
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
	for _, data := range GetData() {
		var times int64
		for _, c := range data.password {
			if c == data.letter {
				times++
			}
		}

		if data.min <= times && times <= data.max {
			valid++
		}
	}

	return utils.TestResult(2, 1, 469, valid)
}

func Part2() int64 {
	var valid int64
	for _, data := range GetData() {
		pwd := []rune(data.password)
		size := int64(len(pwd))
		if data.max <= size && utils.Xor(pwd[data.min-1] == data.letter, pwd[data.max-1] == data.letter) {
			valid++
		}
	}

	return utils.TestResult(2, 2, 267, valid)
}
