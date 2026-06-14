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

One of the interesting things so far is that Go does seem to enjoy a sweet spot for the kinds of problems that AOC throws at people. Go is both fast to compile and run. This means not a lot of time waiting for the code to execute. It's snappy. It also means I'm not worried about having to optimize things so far. I can mostly just trust that the Go compiler will give something fast enough to not worry about speed.

However, on day 11 I've started to see some of the negative sides of Go. First, one of the reasons I'm able to solve these problems fast is because I've done them before. I'm starting to get to some problems that require some thinking and experimentation. This was the first day that I really missed the REPL in Lisp. I've also somewhat missed ipython; it's not as powerful as CL, but it's better than nothing, which is mostly where I find myself with Go.

The other thing I've noticed is Go's verbosity, at least in comparison with Lisp and python. Today's code in Go is 138 LOC. For comparison in python it's 106 LOC and in Lisp it's 71 LOC. I am comparing the v2 of Lisp to this version of Go because both were second efforts at making this code. One of the things that made Lisp so much smaller are the number of external libraries I was able to use, along with dynamic typing. Maybe there are some libraries I could use to make the Go code shorter. But, at no point did I feel like I needed a library to make things less tedious. The code was relatively easy to write, there's just more of it when it comes to LOC.

Having said that, there's another comparison that's interesting. The size in bytes of the Lisp code is 2.62 KB. The size in bytes of the Go code is 2.63 KB. In that sense they are basically identical. This is probably due to Lisp's penchant for long symbol names.

This leads to the following observation: You can do a lot more in one line of Lisp than you can in Go, but each line weighs a lot more.

**Addendum:** I did end up refactoring this day. Now there is far less code repetition. I also added a utility ternary function and made better use function types and functional programming. Still more lines than common Lisp. Still competitive on total bytes.

## [Day 13](day13/day13.go) Mod vs. Rem

This one always trips me up. For this one I pretty much did a translation of the common Lisp code into Go. Most of this day is getting the number theory right, the code itself is pretty trivial.

However, I keep forgetting the difference between modulo and remainder. Common Lisp has a true modulo operator, while Go only has a remainder operator. They give the same answer...until one of the values is negative, and then the remainder gives a negative value. So the straight translation works until you translate mod -> remainder, which is wrong. Took me a while to figure that one out.

## [Day 14](day14/day14.go) Ugly Instructions

At some point Advent of Code starts having ugly instructions. They are well written, but since the process you are applying is itself convoluted, it's easy to misread and debug an algorithm that ends up being the wrong algorithm. Day 14 is when this started.

This is also the first day where I think the Go code compares unfavorably to the Lisp code. The Lisp code is shorter in line count and in bytes. Part of this is the generic power set function I put in utils package. I don't have that in the Go utils package, but there are other functions not needed in Lisp that are in the Go utils package. I think makes them about even in shared code. Having a native bitset also helped out a lot in Lisp.

At this point I should probably start looking into some Go libraries, using [awesome go](https://github.com/avelino/awesome-go) as a resource. While looking up bitset implementations for Go, the author of the library I found was quite proud that it was included there.

## [Day 15](day15/day15.go) Fiddly Algorithm, Just Copied From Previous

I remember this day as having a very simple solution that was difficult to come up with, but simple to code once I had arrived at the algorithm. This meant that implementing it in Go wasn't going to teach me much. So, I just did a translation from the Lisp to Go. I initially tried to remember the algorithm and code it in Go without looking at the previous code, but I ended up making the same mistakes as I did in Lisp when I first implemented it (if I remember correctly).

Still seeing the same trend. Lisp has fewer lines of code, but Go has fewer bytes. This is because:

* Lisp has no syntax, so everything has to be explicitly called out.
* Lisp can't rely on positional conventions to know when/how to do things, again you have to spell it out.
* Because Lisp is by default case-insensitive variables are longer with kebab-casing rather than camel casing.
* Lisp also can't rely on case conventions to export symbols, so you have to manually export them.

So Go programs tend to be shorter, at least in total size, because syntax wins for these types of programs. The promise of Lisp was always that in theory you can build a much higher tower of abstractions than you can in a language like Go. In practice, humans haven't managed to do that, both for lack of knowledge and lack of financial incentive. This means that languages like Go will always do better in the "real world." At least until the "real world" ups it's game.

I think this trend is pretty solid at this point and I don't anticipate commenting on it any further.
