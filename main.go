package main

import (
	"aoc-2020/day01"
	"fmt"
	"os"
	"strconv"
)

type DayFunc func() int64

var all map[int][2]DayFunc = map[int][2]DayFunc{
	1: [2]DayFunc{day01.Part1, day01.Part2},
}

func main() {
	if len(os.Args) == 1 {
		for day, functions := range all {
			fmt.Printf("Day %d 1: %d, 2: %d\n", day, functions[0](), functions[1]())
		}
	} else if len(os.Args) == 2 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		functions := all[day]
		fmt.Printf("Day %d 1: %d, 2: %d\n", day, functions[0](), functions[1]())
	} else if len(os.Args) == 3 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		part, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		fmt.Printf("Day %d %d: %d\n", day, part, all[day][part-1]())
	}
}
