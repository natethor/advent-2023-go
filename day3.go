package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Line struct {
	LineNumber int
	Line       string
}

type Part struct {
	PartNumber int
	LineNumber int
	CharIndex  []int
}

type Symbol struct {
	Symbol     string
	LineNumber int
	CharIndex  int
}

type Gear struct {
	LineNumber int
	CharIndex  int
	PartOne    Part
	PartTwo    Part
}

var linePartsMap = make(map[int][]Part)

func getCharIndex(line string, item string) [][]int {
	var allIndices [][]int

	for i := 0; i <= len(line)-len(item); i++ {
		if line[i:i+len(item)] == item {
			// Check if the characters before and after the found item are symbols or periods
			if (i == 0 || isSymbol(rune(line[i-1]))) && (i+len(item) == len(line) || isSymbol(rune(line[i+len(item)])) || line[i+len(item)] == '.') {
				var indices []int
				for j := i; j < i+len(item); j++ {
					indices = append(indices, j)
				}
				allIndices = append(allIndices, indices)
			}
		}
	}
	return allIndices
}

func extractNumbers(line string) []string {
	return strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
}

func isSymbol(char rune) bool {
	return !unicode.IsLetter(char) && !unicode.IsDigit(char)
}

func isAdjacent(part Part, symbol Symbol) bool {
	// Check if line numbers are adjacent
	if part.LineNumber >= symbol.LineNumber-1 && part.LineNumber <= symbol.LineNumber+1 {
		// Check if symbol is on the same line
		if part.LineNumber == symbol.LineNumber {
			for _, idx := range part.CharIndex {
				// Check if adjacent horizontally
				if idx >= symbol.CharIndex-1 && idx <= symbol.CharIndex+1 {
					return true
				}
			}
		} else {
			// Check if symbol is on the line above or below
			for _, idx := range part.CharIndex {
				// Check if adjacent horizontally or diagonally
				if idx >= symbol.CharIndex-1 && idx <= symbol.CharIndex+1 {
					return true
				}
			}
		}
	}
	return false
}

func findPart(lineNumber, partNumber int) (Part, bool) {
	parts, exists := linePartsMap[lineNumber]
	if !exists {
		return Part{}, false
	}

	for _, p := range parts {
		if p.PartNumber == partNumber {
			return p, true
		}
	}
	return Part{}, false
}

func (l Line) getPartsFromLine() {
	numbers := extractNumbers(l.Line)

	for _, part := range numbers {
		if num, err := strconv.Atoi(part); err == nil {
			if _, exists := findPart(l.LineNumber, num); !exists {
				indicesGroups := getCharIndex(l.Line, part)
				for _, indices := range indicesGroups {
					linePartsMap[l.LineNumber] = append(linePartsMap[l.LineNumber], Part{
						PartNumber: num,
						LineNumber: l.LineNumber,
						CharIndex:  indices,
					})
				}
			}
		}
	}
}

func (l Line) getSymbolsFromLine() []Symbol {
	var symbols []Symbol

	for i, char := range l.Line {
		if isSymbol(char) && char != '.' {
			symbols = append(symbols, Symbol{
				Symbol:     string(char),
				LineNumber: l.LineNumber,
				CharIndex:  i,
			})
		}
	}
	return symbols
}

func (g Gear) getGearRatio() int {
	return g.PartOne.PartNumber * g.PartTwo.PartNumber
}

func partKey(part Part) string {
	return fmt.Sprintf("%d-%d-%v", part.PartNumber, part.LineNumber, part.CharIndex)
}

func findPartsAdjacentToSymbols(parts []Part, symbols []Symbol) []Part {
	uniqueParts := make(map[string]struct{})
	var adjacentParts []Part

	for _, symbol := range symbols {
		for _, part := range parts {
			if isAdjacent(part, symbol) {
				// Check if part is already in the map
				key := partKey(part)
				if _, found := uniqueParts[key]; !found {
					// Add the part to the map and the result slice
					uniqueParts[key] = struct{}{}
					adjacentParts = append(adjacentParts, part)
				}
			}
		}
	}
	return adjacentParts
}

func processSchematic() ([]Part, []Symbol) {
	var parts []Part
	var symbols []Symbol

	if file, err := os.Open("./input/day3.txt"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		parts = make([]Part, 0)
		lineCount := 1
		for scanner.Scan() {
			str := scanner.Text()
			currentLine := Line{LineNumber: lineCount, Line: str}
			currentLine.getPartsFromLine()

			lineParts, exists := linePartsMap[currentLine.LineNumber]
			if exists {
				parts = append(parts, lineParts...)
			}

			symbols = append(symbols, currentLine.getSymbolsFromLine()...)
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	return parts, symbols
}

func Day3Part1() int {
	/*
		add up all the part numbers in the engine schematic
		apparently any number adjacent to a symbol, even diagonally
		is a "part number" and should be included in your sum
		Periods (.) do not count as a symbol.
		What is the sum of all of the part numbers in the engine schematic?
	*/
	parts, symbols := processSchematic()
	valid_parts := findPartsAdjacentToSymbols(parts, symbols)

	// Sum up the valid part numbers
	sum := 0
	for _, part := range valid_parts {
		sum += part.PartNumber
	}
	return sum
}

func Day3Part2() int {
	parts, symbols := processSchematic()

	// loop through symbols finding the * ones and saving them as Gears
	var all_asterisks []Symbol
	for _, symbol := range symbols {
		if symbol.Symbol == "*" {
			all_asterisks = append(all_asterisks, symbol)
		}
	}

	var gears []Gear
	for _, asterisk := range all_asterisks {
		// check if the asterick is adjacent to exactly two Parts
		adjacent_parts := findPartsAdjacentToSymbols(parts, []Symbol{asterisk})
		if len(adjacent_parts) == 2 {
			current_gear := Gear{
				LineNumber: asterisk.LineNumber,
				CharIndex:  asterisk.CharIndex,
				PartOne:    adjacent_parts[0],
				PartTwo:    adjacent_parts[1],
			}
			gears = append(gears, current_gear)
		}

	}

	// Add up the gear ratios of all gears
	var gear_ratios []int
	for _, gear := range gears {
		gear_ratios = append(gear_ratios, gear.getGearRatio())
	}

	sum := 0
	for _, num := range gear_ratios {
		sum += num
	}
	return sum
}
