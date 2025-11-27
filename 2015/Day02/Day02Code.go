package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	length, width, height int
}

func getInput(fileName string) string {
	// Get input from a given file name as one big string
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func convertStringSliceToInt(stringSlice []string) []int {
	// Convert slice of strings into slice of ints
	intSlice := make([]int, 0)
	for _, element := range stringSlice {
		currentVal, _ := strconv.Atoi(element)
		intSlice = append(intSlice, currentVal)
	}
	return intSlice
}

func parseInput() []Box {
	// Parse input as a slice of boxes
	boxes := getInput("boxSizes.txt")
	allBoxes := make([]Box, 0)

	// Loop through each line (each individual box)
	for _, line := range strings.Split(boxes, "\n") {
		// Add box to slice of all boxes
		currentBox := convertStringSliceToInt(strings.Split(line, "x"))
		allBoxes = append(allBoxes, Box{length: currentBox[0],
			width:  currentBox[1],
			height: currentBox[2]})
	}
	return allBoxes
}

func findSmallestSide(box Box) int {
	// Return the smallest side of a box
	side1 := box.length * box.width
	side2 := box.width * box.height
	side3 := box.height * box.length
	if side1 < side2 && side1 < side3 {
		return side1
	} else if side2 < side3 {
		return side2
	}
	return side3
}

func part1() {
	totalWrappingPaper := 0
	// Loop through each box
	for _, box := range parseInput() {
		// 2*l*w + 2*w*h + 2*h*l + smallest side
		totalWrappingPaper += 2 * box.length * box.width
		totalWrappingPaper += 2 * box.width * box.height
		totalWrappingPaper += 2 * box.height * box.length
		totalWrappingPaper += findSmallestSide(box)
	}

	fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalWrappingPaper)
}

func main() {
	part1()
}
