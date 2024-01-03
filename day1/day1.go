package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitStrings = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func getDigit(line string, index int) (int, error) {
	if unicode.IsNumber(rune(line[index])) {
		return strconv.Atoi(string(line[index]))
	}

	fromIndex := line[index:]
	for i, digitStr := range digitStrings {
		if strings.HasPrefix(fromIndex, digitStr) {
			return i, nil
		}
	}

	return -1, nil
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}

	lines := strings.Split(string(content), "\n")

	var sum int
	for _, line := range lines {
		var fDigit, lDigit int = -1, -1
		for i := range line {
			digit, err := getDigit(line, i)
			if err != nil {
				panic(err)
			}
			if digit >= 0 {
				fDigit = digit
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			digit, err := getDigit(line, i)
			if err != nil {
				panic(err)
			}
			if digit >= 0 {
				lDigit = digit
				break
			}
		}

		sum += fDigit*10 + lDigit
	}

	fmt.Println(sum)
}
