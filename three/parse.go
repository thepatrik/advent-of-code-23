package three

import "github.com/thepatrik/advent-of-code-23/pkg/parser"

func parse(filename string) matrix {
	strslice := parser.ReadFile(filename)

	m := make([][]rune, 0)

	for _, line := range strslice {
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		m = append(m, row)
	}

	return m
}
