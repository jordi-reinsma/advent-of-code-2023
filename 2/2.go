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

func PartOne() {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := 0

	for _, line := range strings.Split(input, "\n") {
		sepLine := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(sepLine[0], " ")[1])

		for _, attempt := range strings.Split(sepLine[1], "; ") {
			for _, pick := range strings.Split(attempt, ", ") {
				numberColor := strings.Split(pick, " ")
				number, _ := strconv.Atoi(numberColor[0])
				color := numberColor[1]

				if number > limits[color] {
					goto invalid
				}
			}
		}
		result += gameId
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

		for _, attempt := range strings.Split(strings.Split(line, ": ")[1], "; ") {
			for _, pick := range strings.Split(attempt, ", ") {
				numberColor := strings.Split(pick, " ")
				number, _ := strconv.Atoi(numberColor[0])
				color := numberColor[1]

				if number > minimum[color] {
					minimum[color] = number
				}
			}
		}
		result += minimum["red"] * minimum["green"] * minimum["blue"]
	}
	fmt.Println(result)
}
