package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = []string{"red", "green", "blue"}

func main() {
	content, _ := os.ReadFile("test_input.txt")
	strContent := string(content)
	lines := strings.Split(strContent, "\n")

	var powerSums int
	for _, line := range lines {
		colorMap := make(map[string]int)

		splitStr := strings.Split(line, ":")
		probes := strings.Split(splitStr[1], ";")

		for _, probe := range probes {
			picks := strings.Split(probe, ", ")
			for _, pick := range picks {
				pickNums := strings.Split(strings.Trim(pick, " \t\r\n"), " ")
				number, _ := strconv.Atoi(pickNums[0])
				color := pickNums[1]

				colorMap[color] = max(number, colorMap[color])
			}
		}

		powerProduct := 1
		for _, color := range colors {
			powerProduct *= colorMap[color]
		}
		powerSums += powerProduct
	}

	fmt.Println(powerSums)
}
