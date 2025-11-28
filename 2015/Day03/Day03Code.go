package main

import (
	"fmt"
	"os"
	"strconv"
)

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func moveSanta(character string, x *int, y *int) {
	// Moves x and y coordinates (pass by reference)
	switch character {
	case ">":
		*x++
	case "<":
		*x--
	case "^":
		*y++
	case "v":
		*y--
	}
}

func part1() {
	// Use a map to keep track of visited coords for speed
	x, y := 0, 0
	visitedCoords := make(map[string]bool)
	visitedCoords["0,0"] = true

	for _, character := range getInput("santaDirections.txt") {
		// Move Santa and add new coordinates to map
		moveSanta(string(character), &x, &y)
		visitedCoords[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
	}

	fmt.Printf("Number of houses visited at least once: %d\n", len(visitedCoords))
}

func main() {
	part1()
}
