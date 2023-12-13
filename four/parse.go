package four

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-23/pkg/parser"
)

func parse(filename string) []card {
	strslice := parser.ReadFile(filename)

	cards := make([]card, 0)

	for i := 0; i < len(strslice); i++ {
		linestring := strslice[i]
		card := parseLine(linestring)
		cards = append(cards, card)
	}

	return cards
}

func parseLine(line string) card {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		panic("Invalid input format")
	}

	cardPres := parts[0]

	cardIdS1 := strings.Split(cardPres, " ")
	cardId, err := strconv.Atoi(cardIdS1[len(cardIdS1)-1])
	if err != nil {
		panic(err)
	}

	values := parts[1]

	valuesSplit := strings.Split(values, " | ")
	winningSplit := strings.Split(valuesSplit[0], " ")
	playingSplit := strings.Split(valuesSplit[1], " ")

	winning := make([]int, 0)
	playing := make([]int, 0)

	for _, val := range winningSplit {
		trimmed := strings.TrimSpace(val)
		valInt, err := strconv.Atoi(trimmed)
		if err != nil {
			continue
		}
		winning = append(winning, valInt)
	}

	for _, val := range playingSplit {
		trimmed := strings.TrimSpace(val)
		valInt, err := strconv.Atoi(trimmed)
		if err != nil {
			continue
		}
		playing = append(playing, valInt)
	}

	return card{id: cardId, winning: winning, playing: playing}
}
