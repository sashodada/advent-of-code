package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, length int
}

func inRange(num int, r Range) bool {
	return r.start <= num && num <= r.end()
}

func (r Range) end() int {
	return r.start + r.length - 1
}

func splitRange(target Range, toSplitOver Range) []Range {
	resultRanges := make([]Range, 0)

	rangeStart := target.start
	if toSplitOver.start != target.start && inRange(toSplitOver.start, target) {
		resultRanges = append(resultRanges, Range{rangeStart, toSplitOver.start - rangeStart})
		rangeStart = toSplitOver.start
	}

	if toSplitOver.end() != target.end() && inRange(toSplitOver.end(), target) {
		resultRanges = append(resultRanges, Range{rangeStart, toSplitOver.end() - rangeStart + 1})
		rangeStart = toSplitOver.end() + 1
	}

	resultRanges = append(resultRanges, Range{rangeStart, target.end() - rangeStart + 1})
	return resultRanges
}

func containsRange(target Range, subject Range) bool {
	return inRange(subject.start, target) && inRange(subject.end(), target)
}

func main() {
	fdesc, _ := os.Open("input.txt")
	defer fdesc.Close()

	s := bufio.NewScanner(fdesc)

	s.Scan()
	seedsLine := strings.Trim(strings.Split(s.Text(), ":")[1], " ")
	var seedRanges []Range
	seedRangesArr := strings.Split(seedsLine, " ")
	for i := 0; i < len(seedRangesArr); i += 2 {
		seedRangeStart, _ := strconv.Atoi(seedRangesArr[i])
		seedRangeLen, _ := strconv.Atoi(seedRangesArr[i+1])

		seedRanges = append(seedRanges, Range{seedRangeStart, seedRangeLen})
	}

	unmodifiedSeedRanges := make([]Range, len(seedRanges))
	copy(unmodifiedSeedRanges, seedRanges)
	modifiedSeedRanges := make([]Range, 0)

	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "-") {
			fmt.Println(seedRanges)
			fmt.Println(unmodifiedSeedRanges)
			fmt.Println(modifiedSeedRanges)
			fmt.Println()
			fmt.Println(line)
			seedRanges = append(unmodifiedSeedRanges, modifiedSeedRanges[:]...)
			unmodifiedSeedRanges = make([]Range, len(seedRanges))
			copy(unmodifiedSeedRanges, seedRanges)
			modifiedSeedRanges = make([]Range, 0)
		}
		if line == "" || strings.Contains(line, "-") {
			continue
		}

		rangeInfo := strings.Split(line, " ")
		destRangeStart, _ := strconv.Atoi(rangeInfo[0])
		sourceRangeStart, _ := strconv.Atoi(rangeInfo[1])
		rangeLength, _ := strconv.Atoi(rangeInfo[2])
		sourceRange := Range{sourceRangeStart, rangeLength}

		for i := 0; i < len(unmodifiedSeedRanges); i++ {
			ur := unmodifiedSeedRanges[i]
			if containsRange(sourceRange, ur) {
				modifiedSeedRanges = append(modifiedSeedRanges, Range{ur.start + destRangeStart - sourceRangeStart, ur.length})
				unmodifiedSeedRanges = append(unmodifiedSeedRanges[:i], unmodifiedSeedRanges[i+1:]...)
				i--
			} else if inRange(sourceRange.start, ur) || inRange(sourceRange.end(), ur) {
				unmodifiedSeedRanges = append(unmodifiedSeedRanges, splitRange(ur, sourceRange)[:]...)
				unmodifiedSeedRanges = append(unmodifiedSeedRanges[:i], unmodifiedSeedRanges[i+1:]...)
				i--
			}
		}
	}

	seedRanges = append(unmodifiedSeedRanges, modifiedSeedRanges[:]...)
	minSeedStart := seedRanges[0].start
	for _, r := range seedRanges {
		minSeedStart = min(minSeedStart, r.start)
	}
	fmt.Println(minSeedStart)
}
