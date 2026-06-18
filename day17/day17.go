package day17

import (
	"aoc-2020/utils"
	"slices"
)

func allDiffs(size int, all *[][]int8, soFar []int8) *[][]int8 {
	if len(soFar) == size {
		zeroes := 0
		for i := 0; i < len(soFar); i++ {
			if soFar[i] == 0 {
				zeroes++
			}
		}

		if zeroes != len(soFar) {
			*all = append(*all, slices.Clone(soFar))
		}

		return all
	}

	allDiffs(size, all, append(soFar, int8(-1)))
	allDiffs(size, all, append(soFar, int8(0)))
	return allDiffs(size, all, append(soFar, int8(1)))
}

type coord interface {
	~[3]int8 | ~[4]int8
}

type grid[K coord] struct {
	dimensions int
	data       map[K]bool
	diffs      [][]int8
}

func newGrid[K coord](ksize int) *grid[K] {
	return &grid[K]{
		dimensions: ksize,
		data:       make(map[K]bool),
		diffs:      *allDiffs(ksize, &[][]int8{}, []int8{}),
	}
}

func (g *grid[K]) emptied() *grid[K] {
	return &grid[K]{
		dimensions: g.dimensions,
		data:       make(map[K]bool),
		diffs:      g.diffs,
	}
}

func (g *grid[K]) limits() [][]int8 {
	ret := [][]int8{}

	for dim := 0; dim < g.dimensions; dim++ {
		ret = append(ret, []int8{0, 0})
	}

	for key, _ := range g.data {
		for i := 0; i < g.dimensions; i++ {
			ret[i][0] = min(ret[i][0], key[i])
			ret[i][1] = max(ret[i][1], key[i])
		}
	}

	for i := 0; i < g.dimensions; i++ {
		ret[i][0]--
		ret[i][1]++
	}

	return ret
}

type factory[K comparable] func(x int8, y int8) K
type toKey[K comparable] func(val []int8) K

func parse[K coord](ksize int, fac factory[K]) *grid[K] {
	grid := newGrid[K](ksize)
	for y, str := range utils.ReadLines("inputs/day-17.txt") {
		for x, char := range str {
			if char == '#' {
				grid.data[fac(int8(x), int8(y))] = true
			}
		}
	}

	return grid
}

func add(one []int8, two []int8) []int8 {
	ret := make([]int8, len(one))
	for i := 0; i < len(one); i++ {
		ret[i] = one[i] + two[i]
	}

	return ret
}

func applyConway[K coord](conv toKey[K], old *grid[K], new *grid[K], limits [][]int8, coord []int8) {
	//base case, when len(coords) == new.dimensions, process all neighbors via diffs and adder
	if len(coord) == new.dimensions {
		activeNeighbors := 0
		myKey := conv(coord)
		active := old.data[myKey]
		for _, diff := range new.diffs {
			if old.data[conv(add(coord, diff))] {
				activeNeighbors++
			}
		}

		if active {
			if activeNeighbors == 2 || activeNeighbors == 3 {
				new.data[myKey] = true
			}
		} else {
			if activeNeighbors == 3 {
				new.data[myKey] = true
			}
		}

		return
	}

	//non-base case, loop across limits[0], add to coords, slice off limit
	for i := limits[0][0]; i <= limits[0][1]; i++ {
		applyConway(conv, old, new, limits[1:], append(coord, i))
	}
}

func conway[K coord](conv toKey[K], old *grid[K]) *grid[K] {
	new := old.emptied()
	applyConway(conv, old, new, old.limits(), []int8{})
	return new
}

func Part1() int64 {
	grid := parse(3, func(x int8, y int8) [3]int8 { return [3]int8{x, y, 0} })

	for i := 1; i <= 6; i++ {
		grid = conway(func(a []int8) [3]int8 { return [3]int8{a[0], a[1], a[2]} }, grid)
	}

	return utils.TestResult(17, 1, 213, int64(len(grid.data)))
}

func Part2() int64 {
	grid := parse(4, func(x int8, y int8) [4]int8 { return [4]int8{x, y, 0, 0} })

	for i := 1; i <= 6; i++ {
		grid = conway(func(a []int8) [4]int8 { return [4]int8{a[0], a[1], a[2], a[3]} }, grid)
	}

	return utils.TestResult(17, 1, 1624, int64(len(grid.data)))
}
