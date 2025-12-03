package main

import (
	"fmt"
	"math"
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

func getJoltage(bank string, digitsRemaining int, remainingBankStart int) int {
	bankSlice := stringToIntSlice(bank)
	// Base case: there's just one digit left so return the highest one left
	if digitsRemaining == 1 {
		return slices.Max(bankSlice[remainingBankStart:])
	}

	// Get the next digit and it's next index
	nextDigit := slices.Max(bankSlice[remainingBankStart : len(bank)-digitsRemaining+1])
	remainingBankStart = slices.Index(bankSlice[remainingBankStart:len(bank)-digitsRemaining+1], nextDigit) + remainingBankStart
	// Get current battery joltage (E.g. nextDigit = 9 and digitsRemaining = 5 then currentBatteryJoltage = 90000)
	currentBatteryJoltage := nextDigit * int(math.Pow(10, float64(digitsRemaining-1)))
	return currentBatteryJoltage + getJoltage(bank, digitsRemaining-1, remainingBankStart+1)
}

func part2() {
	totalOutputJoltage := 0
	for bank := range strings.SplitSeq(getInput("banks.txt"), "\n") {
		totalOutputJoltage += getJoltage(strings.TrimSpace(bank), 12, 0)
	}
	fmt.Printf("Total output joltage: %d\n", totalOutputJoltage)
}

func main() {
	part1()
	part2()
}
