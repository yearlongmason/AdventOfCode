package main

import (
	"fmt"
	"os"
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

func main() {
	part1()
}
