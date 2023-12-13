package two

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-23/pkg/parser"
)

func parse(filename string, readIntsAsText bool) []game {
	strslice := parser.ReadFile(filename)

	games := make([]game, 0)

	for i := 0; i < len(strslice); i++ {
		linestring := strslice[i]
		game := parseLine(linestring)
		games = append(games, game)
	}

	return games
}

func parseLine(line string) game {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		panic("Invalid input format")
	}

	values := parts[1]

	subsets := strings.Split(values, "; ")

	sss := make([][]cube, 0)

	for _, subset := range subsets {
		ss := make([]cube, 0)
		cubes := strings.Split(subset, ", ")
		for _, cubestring := range cubes {
			valSplit := strings.Split(cubestring, " ")
			count, _ := strconv.Atoi(strings.TrimSpace(valSplit[0]))

			colorstring := strings.TrimSpace(valSplit[1])

			var c color
			switch colorstring {
			case "red":
				c = red
			case "green":
				c = green
			case "blue":
				c = blue
			default:
				panic("Invalid color")
			}

			ss = append(ss, cube{count: count, color: c})
		}

		sss = append(sss, ss)
	}

	gameID, err := strconv.Atoi(strings.TrimSpace(parts[0][4:]))
	if err != nil {
		panic(err)
	}

	return game{id: gameID, subsets: sss}
}
