package day18

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

var lparen int64 = -1
var rparen int64 = -2
var plus int64 = -3
var mult int64 = -4

func parseLine(s string) []int64 {
	ret := []int64{}
	for _, c := range s {
		switch c {
		case ' ':
			//do nothing
		case '(':
			ret = append(ret, lparen)
		case ')':
			ret = append(ret, rparen)
		case '+':
			ret = append(ret, plus)
		case '*':
			ret = append(ret, mult)
		default:
			ret = append(ret, utils.ToInt64(string(c)))
		}
	}

	return ret
}

func parse() [][]int64 {
	return array.Map(parseLine)(utils.ReadLines("inputs/day-18.txt"))
}

func nextRight(idx int, ops []int64) int {
	num := 1
	for i := idx; i < len(ops); i++ {
		if ops[i] == rparen {
			num--
			if num == 0 {
				return i
			}
		} else if ops[i] == lparen {
			num++
		}
	}

	panic("could not find rparen")
}

type adder func(ops []int64) int64

func solve(ops []int64, f adder) int64 {
	reduced := []int64{}
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case lparen:
			next := nextRight(i+1, ops)
			reduced = append(reduced, solve(ops[(i+1):next], f))
			i = next
		default:
			reduced = append(reduced, ops[i])
		}
	}

	return f(reduced)
}

func inOrder(ops []int64) int64 {
	result := ops[0]
	for i := 1; i < len(ops); i += 2 {
		if ops[i] == plus {
			result += ops[i+1]
		} else {
			result *= ops[i+1]
		}
	}

	return result
}

func Part1() int64 {
	result := int64(0)
	for _, line := range parse() {
		result += solve(line, inOrder)
	}

	return utils.TestResult(18, 1, 11297104473091, result)
}

func plusFirst(ops []int64) int64 {
	reduced := []int64{}
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case plus:
			reduced[len(reduced)-1] += ops[i+1]
			i++
		default:
			reduced = append(reduced, ops[i])
		}
	}

	return inOrder(reduced)
}

func Part2() int64 {
	result := int64(0)
	for _, line := range parse() {
		result += solve(line, plusFirst)
	}

	return utils.TestResult(18, 2, 185348874183674, result)
}
