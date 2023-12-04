package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	PartOne()
	PartTwo()
}

type Game struct {
	id   int
	runs []Run
}

type Run struct {
	subsets []Subset
}

type Subset struct {
	number int
	color  string
}

func parseGame(line string) Game {
	game := Game{}
	sepLine := strings.Split(line, ": ")
	game.id, _ = strconv.Atoi(strings.Split(sepLine[0], " ")[1])
	runs := strings.Split(sepLine[1], "; ")
	game.runs = make([]Run, len(runs))
	for i, run := range runs {
		subsets := strings.Split(run, ", ")
		game.runs[i].subsets = make([]Subset, len(subsets))
		for j, subset := range subsets {
			numberColor := strings.Split(subset, " ")
			number, _ := strconv.Atoi(numberColor[0])
			color := numberColor[1]
			game.runs[i].subsets[j] = Subset{number, color}
		}
	}
	return game
}

func PartOne() {
	result := 0

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		for _, run := range game.runs {
			for _, subset := range run.subsets {
				if subset.number > limits[subset.color] {
					goto invalid
				}
			}
		}
		result += game.id
	invalid:
		continue
	}
	fmt.Println(result)
}

func PartTwo() {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		minimum := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		game := parseGame(line)
		for _, run := range game.runs {
			for _, subset := range run.subsets {
				if subset.number > minimum[subset.color] {
					minimum[subset.color] = subset.number
				}
			}
		}
		result += minimum["red"] * minimum["green"] * minimum["blue"]
	}
	fmt.Println(result)
}
