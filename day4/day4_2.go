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

	var totalCardCount int
	cardsMap := make(map[int]int)
	for _, line := range lines {
		lineElems := strings.Split(line, ":")
		cardId, _ := strconv.Atoi(strings.Trim(strings.ReplaceAll(lineElems[0], "Card ", ""), " "))
		cardsMap[cardId]++
		fmt.Println(cardId, cardsMap[cardId])
		totalCardCount += cardsMap[cardId]

		numbers := strings.Split(lineElems[1], "|")
		winningStrs := strings.Split(strings.Trim(numbers[0], " "), " ")
		currentStrs := strings.Split(strings.Trim(numbers[1], " "), " ")

		winningNums := make(map[int]bool)
		for _, numString := range winningStrs {
			num, _ := strconv.Atoi(numString)
			winningNums[num] = true
		}

		var winningCount int
		for _, numString := range currentStrs {
			num, _ := strconv.Atoi(numString)
			_, isWinningNum := winningNums[num]
			if isWinningNum {
				winningCount++
			}
		}

		for i := 0; i < winningCount; i++ {
			cardsMap[cardId+i+1] += cardsMap[cardId]
		}
	}

	fmt.Println(totalCardCount)
}
