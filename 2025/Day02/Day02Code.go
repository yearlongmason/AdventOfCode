package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type IDRange struct {
	lowerBound, upperBound int
}

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", fileName)
	}

	return string(input)
}

func parseInput() []IDRange {
	IDRanges := make([]IDRange, 0)
	for currentRange := range strings.SplitSeq(getInput("IDRanges.txt"), ",") {
		lowerBound, _ := strconv.Atoi(strings.TrimSpace(strings.Split(currentRange, "-")[0]))
		upperBound, _ := strconv.Atoi(strings.TrimSpace(strings.Split(currentRange, "-")[1]))
		IDRanges = append(IDRanges, IDRange{lowerBound: lowerBound, upperBound: upperBound})
	}

	return IDRanges
}

func isInvalidID(ID string) bool {
	return ID[:len(ID)/2] == ID[len(ID)/2:]
}

func part1() {
	invalidIDSum := 0
	for _, currentRange := range parseInput() {
		// Loop through each possible ID
		for ID := currentRange.lowerBound; ID <= currentRange.upperBound; ID++ {
			if isInvalidID(strconv.Itoa(ID)) {
				invalidIDSum += ID
			}
		}
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", invalidIDSum)
}

func isInvalidIDP2(ID string) bool {
	// Original regular expression: `\b(\d+)\1{1,len(ID)}\b`
	// Unfortunately Go uses the RE2 regular expression engine which does not support backreferences (\1)
	for chunkSize := 1; chunkSize < len(ID); chunkSize++ {
		// Make sure the ID is actually divisible by the chunk size
		if len(ID)%chunkSize != 0 {
			continue
		}

		// Create a pattern that matches the first {chunkSize} characters n times to fill the string
		pattern := `(` + ID[:chunkSize] + `){` + strconv.Itoa(len(ID)/chunkSize) + `}`
		match, _ := regexp.MatchString(pattern, ID)
		if match {
			return true
		}
	}

	return false
}

func part2() {
	invalidIDSum := 0
	for _, currentRange := range parseInput() {
		// Loop through each possible ID
		for ID := currentRange.lowerBound; ID <= currentRange.upperBound; ID++ {
			if isInvalidIDP2(strconv.Itoa(ID)) {
				invalidIDSum += ID
			}
		}
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", invalidIDSum)
}

func main() {
	part1()
	part2()
}
