package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func getNumVowels(str string) int {
	numVowels := strings.Count(str, "a")
	numVowels += strings.Count(str, "e")
	numVowels += strings.Count(str, "i")
	numVowels += strings.Count(str, "o")
	numVowels += strings.Count(str, "u")
	return numVowels
}

func containsDoubleLetter(str string) bool {
	// Loop through each character and if it's the same as the next; return true
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}

	return false
}

func containsIllegalString(str string) bool {
	// Loop through all illegal strings and check if they exist in the given string
	for _, illegalString := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(str, illegalString) {
			return true
		}
	}
	return false
}

func isNice(str string) bool {
	if getNumVowels(str) < 3 {
		return false
	}
	if !containsDoubleLetter(str) {
		return false
	}
	if containsIllegalString(str) {
		return false
	}

	return true
}

func part1() {
	stringsList := getInput("naughtyNice.txt")
	numberOfNiceStrings := 0

	// Loop through each string and if it's nice add 1 to the nice strings
	for str := range strings.SplitSeq(stringsList, "\n") {
		if isNice(str) {
			numberOfNiceStrings++
		}
	}

	fmt.Printf("Number of nice strings: %d\n", numberOfNiceStrings)
}

func main() {
	part1()
}
