package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Could not read file: %s\n", fileName)
	}

	return string(input)
}

func getMD5Hash(key string) string {
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}

func part1() {
	key := getInput("key.txt")
	currentSuffix := 1

	for {
		if strings.HasPrefix(getMD5Hash(key+strconv.Itoa(currentSuffix)), "00000") {
			break
		}
		currentSuffix += 1
	}

	fmt.Printf("First number to produce an MD5 hash with 5 leading 0s: %d\n", currentSuffix)
}

func part2() {
	key := getInput("key.txt")
	currentSuffix := 1

	for {
		if strings.HasPrefix(getMD5Hash(key+strconv.Itoa(currentSuffix)), "000000") {
			break
		}
		currentSuffix += 1
	}

	fmt.Printf("First number to produce an MD5 hash with 6 leading 0s: %d\n", currentSuffix)
}

func main() {
	part1()
	part2()
}
