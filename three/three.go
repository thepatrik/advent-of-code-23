package three

import (
	"errors"
	"strconv"
	"unicode"
)

type matrix [][]rune

func checkPart(r rune) bool {
	if unicode.IsDigit(r) {
		return false
	}
	if r == '.' {
		return false
	}

	return true
}

func checkGear(r rune) bool {
	if unicode.IsDigit(r) {
		return false
	}
	if r == '.' {
		return false
	}

	return true
}

func PartOne(filename string) int {
	m := parse(filename)

	sum := 0
	for y, line := range m {
		currDigit := ""
		isPart := false
		for x, r := range line {
			isDigit := unicode.IsDigit(r)
			if isDigit {
				currDigit += string(r)

				// Look in all 8 directions
				// 1 2 3
				// 4 X 5
				// 6 7 8

				// 1
				if x > 0 && y > 0 {
					if checkPart(m[y-1][x-1]) {
						isPart = true
					}
				}
				// 2
				if y > 0 {
					if checkPart(m[y-1][x]) {
						isPart = true
					}
				}
				// 3
				if x < len(line)-1 && y > 0 {
					if checkPart(m[y-1][x+1]) {
						isPart = true
					}
				}
				// 4
				if x > 0 {
					if checkPart(m[y][x-1]) {
						isPart = true
					}
				}
				// 5
				if x < len(line)-1 {
					if checkPart(m[y][x+1]) {
						isPart = true
					}
				}
				// 6
				if x > 0 && y < len(m)-1 {
					if checkPart(m[y+1][x-1]) {
						isPart = true
					}
				}
				// 7
				if y < len(m)-1 {
					if checkPart(m[y+1][x]) {
						isPart = true
					}
				}
				// 8
				if x < len(line)-1 && y < len(m)-1 {
					if checkPart(m[y+1][x+1]) {
						isPart = true
					}
				}
			}
			if !isDigit || x == len(line)-1 {
				if currDigit != "" {
					if isPart {
						digit, _ := strconv.Atoi(currDigit)
						sum += digit
					}
					currDigit = ""
					isPart = false
				}
			}
		}
	}

	return sum
}

func PartTwo(filename string) int {
	m := parse(filename)

	sum := 0
	for y, line := range m {
		for x, r := range line {
			if r == '*' {
				nums := make([]int, 0)
				// Look up
				if y > 0 {
					upFound := false
					if digit, err := findIntOnLine(m[y-1], x); err == nil {
						nums = append(nums, digit)
						upFound = true
					}
					if !upFound {
						if x > 0 {
							if digit, err := findIntOnLine(m[y-1], x-1); err == nil {
								nums = append(nums, digit)
								upFound = true
							}
						}
						if x < len(line)-1 {
							if digit, err := findIntOnLine(m[y-1], x+1); err == nil {
								nums = append(nums, digit)
								upFound = true
							}
						}
					}
				}
				// Look down
				if y < len(m)-1 {
					downFound := false
					if digit, err := findIntOnLine(m[y+1], x); err == nil {
						nums = append(nums, digit)
						downFound = true
					}
					if !downFound {
						if x > 0 {
							if digit, err := findIntOnLine(m[y+1], x-1); err == nil {
								nums = append(nums, digit)
								downFound = true
							}
						}
						if x < len(line)-1 {
							if digit, err := findIntOnLine(m[y+1], x+1); err == nil {
								nums = append(nums, digit)
								downFound = true
							}
						}
					}
				}

				// Look left
				if x > 0 {
					if digit, err := findIntOnLine(m[y], x-1); err == nil {
						nums = append(nums, digit)
					}
				}

				// Look right
				if x < len(line)-1 {
					if digit, err := findIntOnLine(m[y], x+1); err == nil {
						nums = append(nums, digit)
					}
				}

				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	return sum
}

func findIntOnLine(line []rune, x int) (int, error) {
	if unicode.IsDigit(line[x]) {
		// find first digit index
		start := func() int {
			for i := x; i >= 0; i-- {
				if !unicode.IsDigit(line[i]) {
					return i + 1
				}
			}
			return 0
		}()

		currDigit := ""
		for i := start; i < len(line); i++ {
			if unicode.IsDigit(line[i]) {
				currDigit += string(line[i])
			} else {
				break
			}
		}

		digit, _ := strconv.Atoi(currDigit)
		return digit, nil
	}

	return 0, errors.New("No digit found")
}
