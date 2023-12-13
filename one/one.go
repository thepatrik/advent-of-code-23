package one

import (
	"fmt"
)

var numberMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var errMissingDigit = fmt.Errorf("no digit found in the string")

func PartOne(filename string) int {
	calibrations := parse(filename, false)

	sum := 0
	for _, i := range calibrations {
		sum += i
	}

	return sum
}

func PartTwo(filename string) int {
	calibrations := parse(filename, true)

	sum := 0
	for _, i := range calibrations {
		sum += i
	}

	return sum
}
