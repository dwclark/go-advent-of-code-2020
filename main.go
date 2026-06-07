package main

import (
	"aoc-2020/day01"
	"aoc-2020/day02"
	"aoc-2020/day03"
	"fmt"
	"os"
)

type DayFunc func() int64

var all map[string]map[string]DayFunc = map[string]map[string]DayFunc{
	"1": {"1": day01.Part1, "2": day01.Part2},
	"2": {"1": day02.Part1, "2": day02.Part2},
	"3": {"1": day03.Part1, "2": day03.Part2},
}

func main() {
	if len(os.Args) == 1 {
		for day, functions := range all {
			fmt.Printf("Day %s 1: %d, 2: %d\n", day, functions["1"](), functions["2"]())
		}
	} else if len(os.Args) == 2 {
		functions := all[os.Args[1]]
		fmt.Printf("Day %s 1: %d, 2: %d\n", os.Args[1], functions["1"](), functions["2"]())
	} else if len(os.Args) == 3 {
		fmt.Printf("Day %s %s: %d\n", os.Args[1], os.Args[2], all[os.Args[1]][os.Args[2]]())
	}
}
