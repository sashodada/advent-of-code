package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fdesc, _ := os.Open("input.txt")
	defer fdesc.Close()

	s := bufio.NewScanner(fdesc)

	s.Scan()
	seedsLine := strings.Trim(strings.Split(s.Text(), ":")[1], " ")
	var seedNums []int
	for _, numStr := range strings.Split(seedsLine, " ") {
		seedNum, _ := strconv.Atoi(numStr)
		seedNums = append(seedNums, seedNum)
	}

	newSeedNums := seedNums
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "-") {
			seedNums = newSeedNums
			fmt.Println(seedNums)
			newSeedNums = make([]int, len(seedNums))
			for i, _ := range newSeedNums {
				newSeedNums[i] = -1
			}
		}
		if line == "" || strings.Contains(line, "-") {
			continue
		}

		rangeInfo := strings.Split(line, " ")
		destRangeStart, _ := strconv.Atoi(rangeInfo[0])
		sourceRangeStart, _ := strconv.Atoi(rangeInfo[1])
		rangeLength, _ := strconv.Atoi(rangeInfo[2])

		for i := 0; i < len(seedNums); i++ {
			if seedNums[i] >= sourceRangeStart && seedNums[i] < sourceRangeStart+rangeLength {
				newSeedNums[i] = seedNums[i] + destRangeStart - sourceRangeStart
			} else if newSeedNums[i] == -1 {
				newSeedNums[i] = seedNums[i]
			}
		}
	}

	seedNums = newSeedNums

	sort.Ints(newSeedNums)
	fmt.Println(newSeedNums[0])
}
