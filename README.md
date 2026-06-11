# Advent of Code 2020

## Overview

This repo contains solutions to [Advent of Code 2020](https://adventofcode.com/2020) done in Go. You can build using `make` and then running the resulting binary in build. For example,

To run all of the tests execute:

```build/aoc-2020```

To run the tests for day n:

```build/aoc-2020 n```

To run part p for day n:

```build/aoc-2020 n p```

## Why Am I Doing This Again?

For various reasons I have decided to learn Go. AOC 2020 has always been one of the easier AOCs and so when learning a new language, it's a good choice on which to try new language skills. For the most part doing each day has been good practice for thinking in Go and trying out Go syntax. I don't anticipate commenting much on each day. Go is relatively straightforward to write and read. I have been solving the days in Go without looking at previous solutions. Thus there isn't much insight into how to solve each day as I have already done that before (usually more than once).

## [Day 11](day11/day11.go) Go is Verbose, But Light

One of the interesting things so far is that Go does seem to enjoy a sweet spot for the kinds of problems that AOC throws at people. Go is both fast to compile and run. This means not a lot of time waiting for the code execute. It's snappy. It also means I'm not worried about having to optimize things so far, I can mostly just trust that the Go compiler will give something fast enough to not worry about speed.

However, on day 11 I've started to see some of the negative sides of Go. First, one of the reasons I'm able to solve these problems fast is because I've done them before. I'm starting to get to some problems that require some thinking and experimentation. This was the first day that I really missed the REPL in Lisp. I've also somewhat missed ipython; it's not as powerful as CL, but it's better than nothing, which is mostly where I find myself with Go.

The other thing I've noticed is Go's verbosity, at least in comparison with lisp and python. Today's code in Go is 138 LOC. For comparison in python it's 106 LOC and in lisp it's 71 LOC. I am comparing the v2 of lisp to this version of Go because both were second efforts at making this code. One of the things that made lisp so much smaller are the number of external libraries I was able to use, along with dynamic typing. Maybe there are some libraries I could use to make the Go code shorter. But, at no point did I feel like I needed a library to make things less tedious. The code was relatively easy to write, there's just more of it when it comes to LOC.

Having said that, there's another comparison that's interesting. The size in bytes of the lisp code is 2.62 KB. The size in bytes of the Go code is 2.63 KB. In that sense they are basically identical. This is probably due to lisp's penchant for long symbol names.

This leads to the following observation: You can do a lot more in one line of lisp than you can in Go, but each line weighs a lot more.

**Addendum:** I did end up refactoring this day. Now there is far less code repetition. I also added a utility ternary function and made better use function types and functional programming.
