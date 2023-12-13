package four

import (
	"math"
	"slices"
)

type card struct {
	id      int
	winning []int
	playing []int
}

func PartOne(filename string) int {
	cards := parse(filename)

	sum := 0

	for _, card := range cards {
		matches := 0
		for _, win := range card.winning {
			if slices.Contains(card.playing, win) {
				matches++
			}
		}

		if matches == 1 {
			sum += 1
		} else if matches > 1 {
			sum += int(math.Pow(2, float64(matches-1)))
		}
	}

	return int(sum)
}

func PartTwo(filename string) int {
	cards := parse(filename)

	totalwins := make(map[int]int)

	calculateWins := func(card card, numberOfTimes int) map[int]int {
		winmap := make(map[int]int)
		winmap[card.id] = 1

		wins := 0

		for _, win := range card.winning {
			if slices.Contains(card.playing, win) {
				wins++
			}
		}

		for y := card.id + 1; y <= card.id+wins; y++ {
			winmap[y] += numberOfTimes
		}

		return winmap
	}

	for _, card := range cards {
		numberOfTimes := 1
		val, found := totalwins[card.id]
		if found {
			numberOfTimes += val
		}

		wins := calculateWins(card, numberOfTimes)

		for k, v := range wins {
			totalwins[k] += v
		}
	}

	sum := 0
	for _, v := range totalwins {
		sum += v
	}

	return sum
}
