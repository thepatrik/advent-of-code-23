package two

type color int

const (
	red   color = iota
	green       = iota
	blue        = iota
)

type cube struct {
	count int
	color color
}

type game struct {
	id      int
	subsets [][]cube
}

func PartOne(filename string) int {
	games := parse(filename, false)

	sum := 0

Outer:
	for _, game := range games {
		for _, subset := range game.subsets {
			for _, cube := range subset {
				if cube.color == red && cube.count > 12 {
					continue Outer
				}
				if cube.color == green && cube.count > 13 {
					continue Outer
				}
				if cube.color == blue && cube.count > 14 {
					continue Outer
				}
			}
		}
		sum += game.id
	}

	return sum
}

func PartTwo(filename string) int {
	games := parse(filename, false)

	sum := 0

	for _, game := range games {
		hightestmap := make(map[color]int)

		for _, subset := range game.subsets {
			for _, cube := range subset {
				if hightestmap[cube.color] < cube.count {
					hightestmap[cube.color] = cube.count
				}
			}
		}
		sum += hightestmap[red] * hightestmap[green] * hightestmap[blue]
	}

	return sum
}
