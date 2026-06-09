package day11

import (
	"aoc-2020/utils"
	"fmt"
	//"github.com/IBM/fp-go/array"
)

func seats() [][]rune {
	ret := [][]rune{}
	for _, line := range utils.ReadLines("inputs/day-11.txt") {
		ret = append(ret, []rune(line))
	}

	return ret
}

var empty rune = 'L'
var floor rune = '.'
var occupied rune = '#'
var indexes [3]int = [3]int{-1, 0, 1}

func numOccupied(allSeats [][]rune, row, col int) int {
	numOccupied := 0
	maxRow := len(allSeats)
	maxCol := len(allSeats[0])

	for _, rowDiff := range indexes {
		for _, colDiff := range indexes {
			if rowDiff != 0 || colDiff != 0 {
				testRow := row + rowDiff
				testCol := col + colDiff
				if 0 <= testRow && testRow < maxRow && 0 <= testCol && testCol < maxCol {

					if allSeats[testRow][testCol] == occupied {
						numOccupied++
					}
				}
			}
		}
	}

	return numOccupied
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

func printSeats(seats [][]rune) {
	for _, row := range seats {
		fmt.Println(string(row))
	}

	fmt.Println()
}

func Part1() int64 {
	prev := seats()
	for {
		changes := false
		newSeats := clone(prev)
		for row := 0; row < len(prev); row++ {
			for col := 0; col < len(prev[0]); col++ {
				if prev[row][col] != floor {
					numOccupied := numOccupied(prev, row, col)
					if prev[row][col] == empty && numOccupied == 0 {
						changes = true
						newSeats[row][col] = occupied
					} else if prev[row][col] == occupied && numOccupied >= 4 {
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

	return utils.TestResult(11, 1, 2093, count(occupied, prev))
}

func Part2() int64 {
	return utils.TestResult(11, 2, 0, 0)
}
