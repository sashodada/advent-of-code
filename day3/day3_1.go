package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func hasSymbolNeighbor(matrix [][]rune, row int, startCol int, endCol int) bool {
	for i := max(row-1, 0); i <= min(row+1, len(matrix)-1); i++ {
		for j := max(startCol-1, 0); j <= min(endCol+1, len(matrix[i])-1); j++ {
			char := matrix[i][j]
			if !unicode.IsNumber(char) && char != '.' {
				return true
			}
		}
	}

	return false
}

func main() {
	file, _ := os.ReadFile("input.txt")

	var matrix [][]rune
	for _, line := range strings.Split(string(file), "\r\n") {
		var lineArr []rune
		for _, char := range line {
			lineArr = append(lineArr, char)
		}
		matrix = append(matrix, append(lineArr, '.'))
	}

	var partNumberSum int
	for i, row := range matrix {
		inNumber := false
		var currNumber int
		var firstDigitIndex int

		for j, cell := range row {
			if unicode.IsNumber(cell) {
				if !inNumber {
					firstDigitIndex = j
				}

				inNumber = true
				cellDigit, _ := strconv.Atoi(string(cell))
				currNumber = currNumber*10 + cellDigit
			} else if inNumber {
				if hasSymbolNeighbor(matrix, i, firstDigitIndex, j-1) {
					partNumberSum += currNumber
				}

				currNumber = 0
				inNumber = false
				firstDigitIndex = 0
			}
		}
	}

	fmt.Println(partNumberSum)
}
