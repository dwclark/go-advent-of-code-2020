package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("cannot convert '%s' to an int64", s)
	}

	return int64(num)
}

func ToUInt64(s string) uint64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("cannot convert '%s' to an uint64", s)
	}

	return uint64(num)
}

func ReadLines(name string) []string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return []string{}
	}

	contents := string(bytes)
	trimmed := strings.Trim(contents, "\n\r\t ")
	return strings.Split(trimmed, "\n")
}

func TestResult(day, part, shouldBe, got int64) int64 {
	if shouldBe != got {
		log.Fatalf("For day %d part %d, expected %d, but got %d", day, part, shouldBe, got)
		return -1
	} else {
		return got
	}
}

func Xor(one, two bool) bool {
	var times int
	if one {
		times++
	}

	if two {
		times++
	}

	return times == 1
}

func Tern[T any](condition bool, tval, fval T) T {
	if condition {
		return tval
	} else {
		return fval
	}
}

func Manhattan(vals ...int64) int64 {
	var ret int64
	for _, v := range vals {
		ret += max(v, -v)
	}

	return ret
}
