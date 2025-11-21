package main

import (
	"fmt"
	"os"
	"strconv"
)

func getDigits() string {
	content, err := os.ReadFile("digits.txt")
	if err != nil {
		fmt.Println("Could not find file")
	}

	return string(content)
}

func nextSequence(sequence string) string {
	newSequence := ""
	lastDigit := sequence[0]
	currentLength := 1

	for i := 1; i < len(sequence); i++ {
		// If the digit is the same as the last digit then increment currentLength and move on
		if sequence[i] == lastDigit {
			currentLength++
			continue
		}

		// Otherwise they are different so we need to add to the new sequence and reset lastDigit and currentLength
		newSequence += strconv.Itoa(currentLength) + string(lastDigit)
		lastDigit = sequence[i]
		currentLength = 1
	}

	// Don't forget about the last group
	newSequence += strconv.Itoa(currentLength) + string(lastDigit)

	return newSequence
}

func part1() {
	sequence := getDigits()
	numIterations := 40
	for i := 0; i < numIterations; i++ {
		sequence = nextSequence(sequence)
	}
	fmt.Printf("After %d iterations the length of the sequence is %d\n", numIterations, len(sequence))
}

func part2() {
	sequence := getDigits()
	numIterations := 50
	for i := 0; i < numIterations; i++ {
		sequence = nextSequence(sequence)
	}
	fmt.Printf("After %d iterations the length of the sequence is %d\n", numIterations, len(sequence))
}

func main() {
	part1()
	part2()
}
