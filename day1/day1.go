package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "./input.txt"

func main() {
	maxCals := getMaxCalsFromFile(FILEPATH)
	fmt.Printf("Most calories: %d", maxCals)
}

func getMaxCalsFromFile(filepath string) int {
	normalizedInputBody := getFileAsString(filepath)
	inventories := strings.Split(normalizedInputBody, "\n\n")

	log.Printf("There are %d elves with inventories", len(inventories))
	return getMaxCalories(inventories)

}

// get max calories from a set of inventories
func getMaxCalories(inventories []string) int {
	maxCals := 0
	for elfIdx, inventory := range inventories {
		currCals := getTotalCalories(inventory, elfIdx)
		if currCals > maxCals {
			maxCals = currCals
		}
	}
	return maxCals
}

// get total calories of all food items in an inventory
func getTotalCalories(inventory string, elfIdx int) int {
	total := 0
	foods := strings.Split(inventory, "\n")
	for _, food := range foods {
		if len(food) > 0 {
			foodInt, err := strconv.Atoi(food)
			checkError(err, fmt.Sprintf("Failed Elf#: %d", elfIdx))
			total += foodInt
		}
	}
	return total
}

// read file at filepath, convert to string, then normalize line endings and return the result
func getFileAsString(filepath string) string {
	fileBytes, fileReadErr := os.ReadFile(filepath)
	checkError(fileReadErr, "")
	log.Printf("input.txt successfully read.")

	fullInputBody := string(fileBytes)
	log.Printf("Bytes converted to string.")
	// ensure compatibility on windows and linux
	normalizedInputBody := strings.ReplaceAll(fullInputBody, "\r\n", "\n")
	return normalizedInputBody
}

// If error, panic. If customLog is included, log that first.
func checkError(e error, customLog string) {
	if e != nil {
		if len(customLog) > 0 {
			log.Print(customLog)
		}
		panic(e)
	}
}
