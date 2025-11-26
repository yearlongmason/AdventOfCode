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
	fmt.Printf("The instructions take Santa to floor %d\n", finalFloor)
}

func part2() {
	instructions := getInput("SantaInstructions.txt")
	currentFloor := 0

	// Loop through each instruction
	for index, character := range instructions {
		if string(character) == "(" {
			currentFloor++
		} else {
			currentFloor--
		}

		// Check for floor -1
		if currentFloor == -1 {
			fmt.Printf("The instruction that took Santa to floor -1 was at index %d\n", index+1)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
