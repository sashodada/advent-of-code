package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = []string{"red", "green", "blue"}

func main() {
	content, _ := os.ReadFile("input.txt")

	colorMap := make(map[string]int)
	for i, key := range colors {
		colorMap[key] = 12 + i
	}

	strContent := string(content)
	for _, color := range colors {
		strContent = strings.ReplaceAll(strContent, color, strconv.Itoa(colorMap[color]))
	}

	lines := strings.Split(strContent, "\n")

	var idSums int
	for _, line := range lines {
		isGameValid := true
		splitStr := strings.Split(line, ":")
		game := strings.ReplaceAll(splitStr[0], "Game ", "")
		probes := strings.Split(splitStr[1], ";")

		for _, probe := range probes {
			picks := strings.Split(probe, ", ")
			for _, pick := range picks {
				pickNums := strings.Split(strings.Trim(pick, " \t\r"), " ")
				smaller, _ := strconv.Atoi(pickNums[0])
				larger, _ := strconv.Atoi(pickNums[1])
				if smaller > larger {
					isGameValid = false
				}
			}
		}

		if isGameValid {
			gameId, _ := strconv.Atoi(game)

			idSums += gameId
		}
	}

	fmt.Println(idSums)
}
