package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	PartOne()
	PartTwo()
}

func PartOne() {
	result := 0
	for _, line := range strings.Split(input, "\n") {
		firstDigit, lastDigit := 0, 0
		for _, c := range line {
			if c >= '1' && c <= '9' {
				if firstDigit == 0 {
					firstDigit = int(c - '0')
				}
				lastDigit = int(c - '0')
			}
		}
		result += (firstDigit * 10) + lastDigit
	}
	fmt.Println(result)
}

func PartTwo() {
	digits := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	result := 0
	for _, line := range strings.Split(input, "\n") {
		firstDigit, lastDigit := 0, 0
		firstIdx, lastIdx := len(line), 0

		for i, digit := range digits {
			first := strings.Index(line, digit)
			last := strings.LastIndex(line, digit)
			if first != -1 && first <= firstIdx {
				firstIdx = first
				firstDigit = i%9 + 1
			}
			if last != -1 && last >= lastIdx {
				lastIdx = last
				lastDigit = i%9 + 1
			}
		}
		result += (firstDigit * 10) + lastDigit
	}
	fmt.Println(result)
}
