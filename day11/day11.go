package day11

import (
	"aoc-2020/utils"
	"github.com/IBM/fp-go/array"
)

var empty, floor, occupied, oob rune = 'L', '.', '#', '0'

func seats() [][]rune {
	ret := [][]rune{}
	for _, line := range utils.ReadLines("inputs/day-11.txt") {
		ret = append(ret, []rune(line))
	}

	return ret
}

func element(grid [][]rune, row, col int) rune {
	if row < 0 || col < 0 || len(grid) <= row || len(grid[0]) <= col {
		return oob
	} else {
		return grid[row][col]
	}
}

type neighbor func(grid [][]rune, row, col int) rune

func adj(rowMove, colMove int) neighbor {
	var fun neighbor
	fun = func(grid [][]rune, row, col int) rune {
		return element(grid, row+rowMove, col+colMove)
	}

	return fun
}

func los(rowMove, colMove int) neighbor {
	var fun neighbor
	fun = func(grid [][]rune, row, col int) rune {
		r, c := row+rowMove, col+colMove
		found := element(grid, r, c)
		if found == floor {
			return fun(grid, r, c)
		} else {
			return found
		}
	}

	return fun
}

type change struct {
	row, col int
	to       rune
}

func count(c rune, seats [][]rune) int64 {
	return array.Reduce(func(init int64, row []rune) int64 {
		return array.Reduce(func(init int64, e rune) int64 {
			return utils.Tern(e == c, init+1, init)
		}, init)(row)
	}, 0)(seats)
}

func applyChanges(grid [][]rune, changes []change) bool {
	for _, change := range changes {
		grid[change.row][change.col] = change.to
	}

	return len(changes) > 0
}

func simulate(neighbors []neighbor, limit int) [][]rune {
	all := seats()
	for {
		changes := []change{}
		for row := 0; row < len(all); row++ {
			for col := 0; col < len(all[0]); col++ {
				current := element(all, row, col)
				if current != floor {
					numOccupied := 0
					for _, neighbor := range neighbors {
						if occupied == neighbor(all, row, col) {
							numOccupied++
						}
					}

					if current == empty && numOccupied == 0 {
						changes = append(changes, change{row, col, occupied})
					} else if current == occupied && numOccupied >= limit {
						changes = append(changes, change{row, col, empty})
					}
				}
			}
		}

		if !applyChanges(all, changes) {
			break
		}
	}

	return all
}

func Part1() int64 {
	funcs := []neighbor{adj(-1, -1), adj(-1, 0), adj(-1, 1), adj(0, -1), adj(0, 1),
		adj(1, -1), adj(1, 0), adj(1, 1)}
	return utils.TestResult(11, 1, 2093, count(occupied, simulate(funcs, 4)))
}

func Part2() int64 {
	funcs := []neighbor{los(-1, -1), los(-1, 0), los(-1, 1), los(0, -1), los(0, 1),
		los(1, -1), los(1, 0), los(1, 1)}
	return utils.TestResult(11, 2, 1862, count(occupied, simulate(funcs, 5)))
}
