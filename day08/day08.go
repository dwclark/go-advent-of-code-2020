package day08

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
	"strings"
)

type instruction struct {
	action string
	by     int64
}

func instructions() []instruction {
	f := func(s string) instruction {
		ary := strings.Split(s, " ")
		return instruction{ary[0], utils.ToInt64(ary[1])}
	}

	return array.Map(f)(utils.ReadLines("inputs/day-08.txt"))
}

func run(instructions []instruction) (int64, bool) {
	pointer := 0
	visited := map[int]bool{}
	var accumulator int64
	for {
		if pointer == len(instructions) {
			return accumulator, true
		}

		if visited[pointer] {
			return accumulator, false
		}

		visited[pointer] = true
		switch i := instructions[pointer]; i.action {
		case "nop":
			pointer++
		case "acc":
			pointer++
			accumulator += i.by
		case "jmp":
			pointer += int(i.by)
		}
	}
}

func Part1() int64 {
	accumulator, _ := run(instructions())
	return utils.TestResult(8, 1, 1782, accumulator)
}

func swap(i instruction) instruction {
	if i.action == "nop" {
		return instruction{"jmp", i.by}
	} else {
		return instruction{"nop", i.by}
	}
}

func Part2() int64 {
	instructions := instructions()
	var accumulator int64
	var correct bool

	for idx := 0; idx < len(instructions); idx++ {
		if instructions[idx].action == "acc" {
			continue
		}

		instructions[idx] = swap(instructions[idx])
		accumulator, correct = run(instructions)
		if correct {
			break
		} else {
			instructions[idx] = swap(instructions[idx])
		}
	}

	return utils.TestResult(8, 2, 797, accumulator)
}
