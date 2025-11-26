package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(fileName string) string {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file: %s\n", fileName)
		return ""
	}

	return string(fileContents)
}

func part1() {
	instructions := getInput("SantaInstructions.txt")
	finalFloor := strings.Count(instructions, "(") - strings.Count(instructions, ")")
	fmt.Printf("The instructions take Santa to floor %d", finalFloor)
}

func main() {
	part1()
}
