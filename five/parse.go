package five

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-23/pkg/parser"
)

func parse(filename string) almanac {
	strslice := parser.ReadFile(filename)

	first := strslice[0]
	seeds := strings.Split(first, "seeds: ")
	seedSlice := strings.Split(seeds[1], " ")
	seedInts := make([]int, 0)
	for _, seed := range seedSlice {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seedInts = append(seedInts, seedInt)
	}

	maps := make([]mapstruct, 0)
	ranges := make([]valueRange, 0)

	for i := 3; i < len(strslice); i++ {
		linestring := strslice[i]
		if linestring == "" {
			continue
		}
		if strings.Contains(linestring, " map:") {
			m := mapstruct{ranges: ranges}
			maps = append(maps, m)
			ranges = make([]valueRange, 0)
			continue
		}

		v := func() valueRange {
			v := valueRange{}
			intsplit := strings.Split(linestring, " ")
			for i, str := range intsplit {
				num, _ := strconv.Atoi(str)
				switch i {
				case 0:
					v.destination = num
				case 1:
					v.source = num
				case 2:
					v.jumps = num
				}
			}
			return v
		}()

		ranges = append(ranges, v)
	}

	m := mapstruct{ranges: ranges}
	maps = append(maps, m)

	return almanac{seeds: seedInts, maps: maps}
}
