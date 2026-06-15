package day13

import (
	"aoc-2020/utils"
	"math"
	"strings"
)

type bus struct {
	index, id int64
}

func parse() (int64, []bus) {
	lines := utils.ReadLines("inputs/day-13.txt")
	buses := []bus{}
	for idx, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			buses = append(buses, bus{int64(idx), utils.ToInt64(id)})
		}
	}

	return utils.ToInt64(lines[0]), buses
}

func Part1() int64 {
	earliest, buses := parse()
	var val int64

outer:
	for i := earliest + 1; i < math.MaxInt64; i++ {
		for _, bus := range buses {
			if i%bus.id == 0 {
				val = bus.id * (i - earliest)
				break outer
			}
		}
	}

	return utils.TestResult(13, 1, 205, val)
}

func (b bus) remainder() int64 {
	rem := (b.id - b.index) % b.id
	if rem >= 0 {
		return rem
	} else {
		return b.id + rem
	}
}

func Part2() int64 {
	_, buses := parse()
	result := int64(0)
	increment := int64(1)

	for _, bus := range buses {
		rem := bus.remainder()
		for result%bus.id != rem {
			result += increment
		}

		increment = increment * bus.id
	}

	return utils.TestResult(13, 2, 803025030761664, result)
}
