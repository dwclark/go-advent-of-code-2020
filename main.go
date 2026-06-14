package main

import (
	"aoc-2020/day01"
	"aoc-2020/day02"
	"aoc-2020/day03"
	"aoc-2020/day04"
	"aoc-2020/day05"
	"aoc-2020/day06"
	"aoc-2020/day07"
	"aoc-2020/day08"
	"aoc-2020/day09"
	"aoc-2020/day10"
	"aoc-2020/day11"
	"aoc-2020/day12"
	"aoc-2020/day13"
	"aoc-2020/day14"
	"aoc-2020/day15"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DayFunc func() int64

var all [][]DayFunc = [][]DayFunc{
	{day01.Part1, day01.Part2},
	{day02.Part1, day02.Part2},
	{day03.Part1, day03.Part2},
	{day04.Part1, day04.Part2},
	{day05.Part1, day05.Part2},
	{day06.Part1, day06.Part2},
	{day07.Part1, day07.Part2},
	{day08.Part1, day08.Part2},
	{day09.Part1, day09.Part2},
	{day10.Part1, day10.Part2},
	{day11.Part1, day11.Part2},
	{day12.Part1, day12.Part2},
	{day13.Part1, day13.Part2},
	{day14.Part1, day14.Part2},
	{day15.Part1, day15.Part2},
}

func main() {
	if len(os.Args) == 1 {
		for idx, functions := range all {
			day := idx + 1
			fmt.Printf("Day %d 1: %d, 2: %d\n", day, functions[0](), functions[1]())
		}
	} else if len(os.Args) == 2 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("could not convert %s to number", os.Args[1])
		}

		functions := all[day-1]
		fmt.Printf("Day %d 1: %d, 2: %d\n", day, functions[0](), functions[1]())
	} else if len(os.Args) == 3 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("could not convert %s to number", os.Args[1])
		}

		part, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("could not convert %s to number", os.Args[2])
		}

		fmt.Printf("Day %d %d: %d\n", day, part, all[day-1][part-1]())
	}
}
