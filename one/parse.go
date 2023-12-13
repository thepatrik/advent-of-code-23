package one

import (
	"strings"
	"unicode"

	"github.com/thepatrik/advent-of-code-23/pkg/parser"
)

func parse(filename string, readIntsAsText bool) []int {
	strslice := parser.ReadFile(filename)

	calibrations := make([]int, 0)

	for i := 0; i < len(strslice); i++ {
		linestring := strslice[i]
		first, _ := firstDigit(linestring, readIntsAsText)
		last, _ := lastDigit(linestring, readIntsAsText)

		calibrations = append(calibrations, first*10+last)
	}

	return calibrations
}

func firstDigit(input string, parseIntFromString bool) (int, error) {
	runes := []rune(input)

	if value, found := numberMap[input]; found {
		return value, nil
	}

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsDigit(r) {
			return int(r - '0'), nil
		}

		if !parseIntFromString {
			continue
		}

		if val, err := intInString(input[i:]); err == nil {
			return val, nil
		}
	}

	return 0, errMissingDigit
}

func lastDigit(input string, parseIntFromString bool) (int, error) {
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsDigit(r) {
			return int(r - '0'), nil
		}

		if !parseIntFromString {
			continue
		}

		if val, err := intInString(input[i:]); err == nil {
			return val, nil
		}
	}

	return 0, errMissingDigit
}

func intInString(s string) (int, error) {
	for k, v := range numberMap {
		if strings.HasPrefix(s, k) {
			return v, nil
		}
	}

	return 0, errMissingDigit
}
