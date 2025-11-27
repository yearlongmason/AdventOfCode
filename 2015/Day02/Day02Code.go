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
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func convertStringSliceToInt(stringSlice []string) []int {
	intSlice := make([]int, 0)
	for _, element := range stringSlice {
		currentVal, _ := strconv.Atoi(element)
		intSlice = append(intSlice, currentVal)
	}
	return intSlice
}

func parseInput() []Box {
	boxes := getInput("boxSizes.txt")
	allBoxes := make([]Box, 0)

	currentBox := make([]int, 0)
	for index, line := range strings.Split(boxes, "\n") {
		currentBox = strings.Split(line, "x")
		fmt.Println(index, strings.Split(currentBox, "x"))
	}
	allBoxes = append(allBoxes, Box{1, 2, 3})
	return allBoxes
}

func part1() {
	fmt.Println(parseInput()[0])
}

func main() {
	part1()
}
