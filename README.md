# Advent of Code 2020

## Overview

This repo contains solutions to [Advent of Code 2020](https://adventofcode.com/2020) done in Go. You can build using `make` and then running the resulting binary in build. For example,

To run all of the tests execute:

```build/aoc-2020```

To run the tests for day n:

```build/aoc-2020 n```

To run part p for day n:

```build/aoc-2020 n p```

## Code Organization

To prevent name collisions I decided to put each day in its own package. Thus you can find day 1 in [day01/day01.go](day01/day01.go), and so forth. Each package has the functions `Part1` and `Part2` so that they can be called from the main package.

## Why Am I Doing This Again?

For various reasons I have decided to learn Go. AOC 2020 has always been one of the easier AOCs and so when learning a new language, it's a good choice on which to try new language skills. For the most part doing each day has been good practice for thinking in Go and trying out Go syntax. I don't anticipate commenting much on each day. Go is relatively straightforward to write and read. I have been solving the days in Go without looking at previous solutions. Thus there isn't much insight into how to solve each day as I have already done that before (usually more than once).

