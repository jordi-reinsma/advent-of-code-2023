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
	result := 0

	lines := strings.Split(input, "\n")
	engine := make([][]byte, len(lines))
	for i, line := range lines {
		engine[i] = []byte(line)
	}

	for i := 0; i < len(engine); i++ {
		number := strings.Builder{}
		for j := 0; j < len(engine[i]); j++ {
			for ; j < len(engine[i]) && isNumber(engine[i][j]); j++ {
				number.WriteByte(engine[i][j])
			}
			if number.Len() > 0 {
				symbols := getAdjacentSymbols(engine, isAnySymbol, i-1, j-number.Len()-1, i+1, j)
				result += toNumber(number) * len(symbols)
				number.Reset()
			}
		}
	}

	fmt.Println(result)
}

func PartTwo() {
	result := 0

	lines := strings.Split(input, "\n")
	engine := make([][]byte, len(lines))
	for i, line := range lines {
		engine[i] = []byte(line)
	}

	partsInGears := make(map[Point][]int)

	for i := 0; i < len(engine); i++ {
		number := strings.Builder{}
		for j := 0; j < len(engine[i]); j++ {
			for ; j < len(engine[i]) && isNumber(engine[i][j]); j++ {
				number.WriteByte(engine[i][j])
			}
			if number.Len() > 0 {
				symbols := getAdjacentSymbols(engine, isGear, i-1, j-number.Len()-1, i+1, j)
				for _, symbol := range symbols {
					partsInGears[symbol] = append(partsInGears[symbol], toNumber(number))
				}
				number.Reset()
			}
		}
	}
	for _, parts := range partsInGears {
		if len(parts) == 2 {
			result += parts[0] * parts[1]
		}
	}

	fmt.Println(result)
}

type Point struct {
	x, y int
}

func isGear(c byte) bool {
	return c == '*'
}
func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAnySymbol(c byte) bool {
	return !isNumber(c) && c != '.'
}

func toNumber(b strings.Builder) int {
	number, _ := strconv.Atoi(b.String())
	return number
}

func getAdjacentSymbols(engine [][]byte, findSymbol func(byte) bool, i, j, k, l int) []Point {
	i, j = max(i, 0), max(j, 0)
	k, l = min(k, len(engine)-1), min(l, len(engine[0])-1)

	symbols := make([]Point, 0)
	for x := i; x <= k; x++ {
		for y := j; y <= l; y++ {
			if findSymbol(engine[x][y]) {
				symbols = append(symbols, Point{x, y})
			}
		}
	}
	return symbols
}
