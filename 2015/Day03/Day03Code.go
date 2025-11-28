package main

import (
	"fmt"
	"os"
)

type Coordinate struct {
	x, y int
}

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func main() {
	fmt.Println(getInput("santaDirections.txt"))
}
