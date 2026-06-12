package day12

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

type instruction struct {
	action rune
	val    int64
}

func parse(s string) instruction {
	return instruction{[]rune(s)[0], utils.ToInt64(s[1:])}
}

func instructions() []instruction {
	return array.Map(parse)(utils.ReadLines("inputs/day-12.txt"))
}

func Part1() int64 {
	x, y, dir := int64(0), int64(0), int64(90)
	for _, instr := range instructions() {
		switch instr.action {
		case 'N':
			y += instr.val
		case 'S':
			y -= instr.val
		case 'E':
			x += instr.val
		case 'W':
			x -= instr.val
		case 'L':
			dir = (dir + (360 - instr.val)) % 360
		case 'R':
			dir = (dir + instr.val) % 360
		default: //forward
			switch dir {
			case 0:
				y += instr.val
			case 90:
				x += instr.val
			case 180:
				y -= instr.val
			default: //270
				x -= instr.val
			}
		}
	}

	return utils.TestResult(12, 1, 445, (utils.Manhattan(x, y)))
}

func Part2() int64 {
	x, y, wx, wy := int64(0), int64(0), int64(10), int64(1)
	for _, instr := range instructions() {
		switch instr.action {
		case 'N':
			wy += instr.val
		case 'S':
			wy -= instr.val
		case 'E':
			wx += instr.val
		case 'W':
			wx -= instr.val
		case 'L':
			switch instr.val {
			case 90:
				wx, wy = -wy, wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = wy, -wx
			}
		case 'R':
			switch instr.val {
			case 90:
				wx, wy = wy, -wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = -wy, wx
			}
		default: //forward
			x, y = x+(wx*instr.val), y+(wy*instr.val)
		}
	}

	return utils.TestResult(12, 2, 42495, (utils.Manhattan(x, y)))
}
