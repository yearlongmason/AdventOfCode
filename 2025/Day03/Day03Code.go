package main

import (
	"fmt"
	"os"
)

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", fileName)
	}

	return string(input)
}

func part1() {
	fmt.Println(getInput("banks.txt"))
}

func main() {
	part1()
}
