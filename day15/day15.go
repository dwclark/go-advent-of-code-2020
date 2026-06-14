package day15

import (
	"aoc-2020/utils"
)

var plays []int = []int{2, 20, 0, 4, 1, 17}
var board map[int]int

func speak(num int, turn int) int {
	prevTurn, isThere := board[num]
	board[num] = turn
	if isThere {
		return turn - prevTurn
	} else {
		return 0
	}
}

func play(maxTurns int) int {
	//reset the board
	board = map[int]int{}
	for i, play := range plays {
		speak(play, i)
	}

	//play the game
	prev := 0
	for turn := len(plays); turn < (maxTurns - 1); turn++ {
		prev = speak(prev, turn)
	}

	return prev
}

func Part1() int64 {
	return utils.TestResult(15, 1, 758, int64(play(2020)))
}

func Part2() int64 {
	return utils.TestResult(15, 1, 814, int64(play(30_000_000)))
}
