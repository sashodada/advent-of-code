package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(strings.ReplaceAll(string(file), "  ", " "), "\r\n")

	var totalScore int
	for _, line := range lines {
		numbers := strings.Split(line, "|")
		winningStrs := strings.Split(strings.Trim(numbers[0], " "), " ")
		currentStrs := strings.Split(strings.Trim(numbers[1], " "), " ")

		winningNums := make(map[int]bool)
		for _, numString := range winningStrs {
			num, _ := strconv.Atoi(numString)
			winningNums[num] = true
		}

		var score int
		for _, numString := range currentStrs {
			num, _ := strconv.Atoi(numString)
			_, isWinningNum := winningNums[num]
			if isWinningNum {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		totalScore += score
	}

	fmt.Println(totalScore)
}
