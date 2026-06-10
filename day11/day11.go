package day11

import (
	"aoc-2020/utils"
)

func seats() [][]rune {
	ret := [][]rune{}
	for _, line := range utils.ReadLines("inputs/day-11.txt") {
		ret = append(ret, []rune(line))
	}

	return ret
}

type move struct {
	row, col int
}

type computeOccupied func(all [][]rune, row, col int) int

var empty rune = 'L'
var floor rune = '.'
var occupied rune = '#'
var moves [8]move = [8]move{move{-1, -1}, move{-1, 0}, move{-1, 1}, move{0, -1}, move{0, 1},
	move{1, -1}, move{1, 0}, move{1, 1}}

func adjacent(allSeats [][]rune, row, col int) int {
	ret := 0
	maxRow := len(allSeats)
	maxCol := len(allSeats[0])

	for _, move := range moves {
		testRow := row + move.row
		testCol := col + move.col
		if 0 <= testRow && testRow < maxRow && 0 <= testCol && testCol < maxCol {

			if allSeats[testRow][testCol] == occupied {
				ret++
			}
		}
	}

	return ret
}

func lineOfSight(allSeats [][]rune, row, col int) int {
	ret := 0
	maxRow := len(allSeats)
	maxCol := len(allSeats[0])

	for _, move := range moves {
		multiplier := 1
		for {
			testRow := row + (multiplier * move.row)
			testCol := col + (multiplier * move.col)
			if 0 <= testRow && testRow < maxRow && 0 <= testCol && testCol < maxCol {
				toTest := allSeats[testRow][testCol]
				if toTest == occupied {
					ret++
					break
				} else if toTest == empty {
					break
				} else {
					multiplier++
				}
			} else {
				break
			}
		}
	}

	return ret
}

func clone(seats [][]rune) [][]rune {
	ret := [][]rune{}

	for _, theSlice := range seats {
		tmp := make([]rune, len(theSlice))
		copy(tmp, theSlice)
		ret = append(ret, tmp)
	}

	return ret
}

func count(c rune, seats [][]rune) int64 {
	var numEmpty int64
	for row := 0; row < len(seats); row++ {
		for col := 0; col < len(seats[0]); col++ {
			if seats[row][col] == c {
				numEmpty++
			}
		}
	}

	return numEmpty
}

func simulate(comp computeOccupied, limit int) [][]rune {
	prev := seats()
	for {
		changes := false
		newSeats := clone(prev)
		for row := 0; row < len(prev); row++ {
			for col := 0; col < len(prev[0]); col++ {
				if prev[row][col] != floor {
					numOccupied := comp(prev, row, col)
					if prev[row][col] == empty && numOccupied == 0 {
						changes = true
						newSeats[row][col] = occupied
					} else if prev[row][col] == occupied && numOccupied >= limit {
						changes = true
						newSeats[row][col] = empty
					}
				}
			}
		}

		prev = newSeats
		if !changes {
			break
		} else {
			changes = false
		}
	}

	return prev
}

func Part1() int64 {
	return utils.TestResult(11, 1, 2093, count(occupied, simulate(adjacent, 4)))
}

func Part2() int64 {
	return utils.TestResult(11, 2, 1862, count(occupied, simulate(lineOfSight, 5)))
}
