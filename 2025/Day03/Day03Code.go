package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", fileName)
	}

	return string(input)
}

func stringToIntSlice(str string) []int {
	// Convert bank (string of digits) into a slice of ints
	intSlice := make([]int, 0)
	for _, digit := range str {
		digit, err := strconv.Atoi(strings.TrimSpace(string(digit)))
		if err == nil {
			intSlice = append(intSlice, digit)
		}
	}
	return intSlice
}

func part1() {
	totalOutputJoltage := 0
	for bank := range strings.SplitSeq(getInput("banks.txt"), "\n") {
		bankSlice := stringToIntSlice(bank)
		// Find first and second digit and add to total output joltage
		firstDigit := slices.Max(bankSlice[:len(bankSlice)-1])
		secondDigit := slices.Max(bankSlice[slices.Index(bankSlice, firstDigit)+1:])
		totalOutputJoltage += (firstDigit * 10) + secondDigit
	}

	fmt.Printf("Total output joltage: %d\n", totalOutputJoltage)
}

func main() {
	part1()
}
