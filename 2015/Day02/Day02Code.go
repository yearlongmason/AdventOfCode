package main

import (
	"fmt"
	"os"
	"slices"
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
	// Helper function that convert slice of strings into slice of ints
	intSlice := make([]int, 0)

	// Cast each element in the slice as an int and add to intSlice
	for _, element := range stringSlice {
		currentVal, err := strconv.Atoi(element)
		// Sometimes casting as int doesn't work because of '\n' at the end
		// If that happens, try again but without the last character
		if err != nil {
			currentVal, err = strconv.Atoi(element[0 : len(element)-1])
		}
		intSlice = append(intSlice, currentVal)
	}
	return intSlice
}

func parseInput() []Box {
	// Parse input as a slice of boxes
	allBoxes := make([]Box, 0)

	// Loop through each line of the input (each individual box)
	for _, line := range strings.Split(getInput("boxSizes.txt"), "\n") {
		// Add box to slice of all boxes
		currentBox := convertStringSliceToInt(strings.Split(line, "x"))
		allBoxes = append(allBoxes, Box{length: currentBox[0],
			width:  currentBox[1],
			height: currentBox[2]})
	}
	return allBoxes
}

func part1() {
	totalWrappingPaper := 0
	// Loop through each box
	for _, box := range parseInput() {
		// 2*l*w + 2*w*h + 2*h*l + smallest side
		totalWrappingPaper += 2 * box.length * box.width
		totalWrappingPaper += 2 * box.width * box.height
		totalWrappingPaper += 2 * box.height * box.length
		totalWrappingPaper += slices.Min([]int{box.length * box.width,
			box.width * box.height,
			box.height * box.length})
	}

	fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalWrappingPaper)
}

func getRibbonLength(box Box) int {
	ribbonLength := 0

	// Getting perimeter of two smallest sides by adding all sides and subtracting biggest side
	ribbonLength += box.length * 2
	ribbonLength += box.height * 2
	ribbonLength += box.width * 2
	ribbonLength -= slices.Max([]int{box.length, box.width, box.height}) * 2

	// Don't forget the bow!
	ribbonLength += box.length * box.height * box.width
	return ribbonLength
}

func part2() {
	totalRibbon := 0

	for _, box := range parseInput() {
		totalRibbon += getRibbonLength(box)
	}
	fmt.Printf("Total feet of ribbon needed: %d\n", totalRibbon)
}

func main() {
	part1()
	part2()
}
