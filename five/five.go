package five

import (
	"math"
)

type valueRange struct {
	source      int
	destination int
	jumps       int
}

type mapstruct struct {
	ranges []valueRange
}

type almanac struct {
	seeds []int
	maps  []mapstruct
}

func PartOne(filename string) int {
	almanac := parse(filename)

	findDestination := func(i int, m mapstruct) int {
		for _, r := range m.ranges {
			if i >= r.source && i < r.source+r.jumps {
				return i - r.source + r.destination
			}
		}
		return i
	}

	lowest := math.MaxInt
	for _, seed := range almanac.seeds {
		for _, m := range almanac.maps {
			seed = findDestination(seed, m)
		}
		if seed < lowest {
			lowest = seed
		}
	}

	return lowest
}

func PartTwo(filename string) int {
	almanac := parse(filename)

	isSeed := func(val int) bool {
		for i := 0; i < len(almanac.seeds); i += 2 {
			initialSeed := almanac.seeds[i]
			rangeVal := almanac.seeds[i+1]

			if val >= initialSeed && val <= initialSeed+rangeVal {
				return true
			}
		}
		return false
	}

	findSource := func(i int, m mapstruct) int {
		for _, r := range m.ranges {
			if i >= r.destination && i < r.destination+r.jumps {
				return i - r.destination + r.source
			}
		}
		return i
	}

	// Instead of looping through all possible seed values, loop through possible
	// location values in ascending order and check if they are seeds. When a seed
	// is found, the job is done.
	seedf := func() int {
		for location := 0; location < math.MaxInt; location++ {
			val := location
			for i := len(almanac.maps) - 1; i >= 0; i-- {
				val = findSource(val, almanac.maps[i])
			}

			if isSeed(val) {
				return location
			}
		}

		panic("no seed found")
	}

	return seedf()
}
