package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	PartOne()
	PartTwo()
}

type Scratchcard struct {
	winners []int
	picked  []int
}

func parseScratchcard(line string) Scratchcard {
	scratchcard := Scratchcard{}
	numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
	winners := strings.Fields(numbers[0])
	picked := strings.Fields(numbers[1])

	scratchcard.winners = make([]int, len(winners))
	for i, winner := range winners {
		scratchcard.winners[i], _ = strconv.Atoi(winner)
	}

	scratchcard.picked = make([]int, len(picked))
	for i, pick := range picked {
		scratchcard.picked[i], _ = strconv.Atoi(pick)
	}
	return scratchcard
}

func calcWins(scratchcard Scratchcard) int {
	wins := 0
	for _, picked := range scratchcard.picked {
		if slices.Contains(scratchcard.winners, picked) {
			wins++
		}
	}
	return wins
}

func calcScore(wins int) int {
	if wins <= 1 {
		return wins
	} else {
		return 1 << (wins - 1)
	}
}

func PartOne() {
	result := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		scratchcard := parseScratchcard(line)
		wins := calcWins(scratchcard)
		result += calcScore(wins)
	}
	fmt.Println(result)
}

func PartTwo() {
	result := 0

	lines := strings.Split(input, "\n")
	winList := make([]int, len(lines))
	copies := make([]int, len(lines))
	for i, line := range lines {
		winList[i] = calcWins(parseScratchcard(line))
		copies[i] = 1
	}

	for i := 0; i < len(winList); i++ {
		wins := winList[i]
		for j := i + 1; j <= i+wins && j < len(winList); j++ {
			copies[j] += copies[i]
		}
		result += copies[i]
	}

	fmt.Println(result)
}
