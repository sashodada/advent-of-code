package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getCoordKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func appendPartToStarsMap(matrix [][]rune, row int, startCol int, endCol int, starsMap map[string][]int, partNumber int) {
	for i := max(row-1, 0); i <= min(row+1, len(matrix)-1); i++ {
		for j := max(startCol-1, 0); j <= min(endCol+1, len(matrix[i])-1); j++ {
			if matrix[i][j] == '*' {
				key := getCoordKey(i, j)
				starsMap[key] = append(starsMap[key], partNumber)
			}
		}
	}
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

	starsMap := make(map[string][]int)
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
				appendPartToStarsMap(matrix, i, firstDigitIndex, j-1, starsMap, currNumber)

				currNumber = 0
				inNumber = false
				firstDigitIndex = 0
			}
		}
	}

	var partRatioSum int
	for _, a := range starsMap {
		if len(a) == 2 {
			partRatioSum += a[0] * a[1]
		}
	}

	fmt.Println(partRatioSum)
}
