package main

import (
	"fmt"
	"os"
)

func getInput(fileName string) string {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file: %s\n", fileName)
		return ""
	}

	return string(fileContents)
}

func main() {
	fmt.Println(getInput("SantaInstructions.txt"))
}
