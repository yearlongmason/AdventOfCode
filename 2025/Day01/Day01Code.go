package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	direction string
	rotations int
}

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func getRotations() []int {
	// Create a slice of ints from the input file that get the number of rotations
	// Positive number if the direction is right, negative if the direction is left
	instructions := make([]int, 0)
	for currentInstruction := range strings.SplitSeq(getInput("instructions.txt"), "\n") {
		direction := string(currentInstruction[0])
		rotations, _ := strconv.Atoi(strings.TrimSpace(currentInstruction[1:]))
		if direction == "R" {
			instructions = append(instructions, rotations)
		} else {
			instructions = append(instructions, rotations*-1)
		}
	}
	return instructions
}

func part1() {
	numZeros := 0
	dial := 50
	// Loop through each rotation
	for _, rotations := range getRotations() {
		// Move the dial and account for overflow/underflow
		dial += rotations
		dial %= 100
		if dial == 0 {
			numZeros++
		}
	}

	fmt.Printf("Number of times the dial lands on 0: %d\n", numZeros)
}

func getInstructions() []Instruction {
	// Create a slice of Instructions from the input file
	instructions := make([]Instruction, 0)
	for currentInstruction := range strings.SplitSeq(getInput("instructions.txt"), "\n") {
		direction := string(currentInstruction[0])
		rotations, _ := strconv.Atoi(strings.TrimSpace(currentInstruction[1:]))
		instructions = append(instructions, Instruction{direction: direction, rotations: rotations})
	}
	return instructions
}

func part2() {

}

func main() {
	start := time.Now()
	part1()
	part2()
	fmt.Printf("Runtime: %f seconds", time.Since(start).Seconds())
}
