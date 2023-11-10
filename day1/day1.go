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
	top3 := getTop3FromFile(FILEPATH)
	solvePuzzle1(top3)
	solvePuzzle2(top3)
}

func getTop3FromFile(filepath string) [3]int {
	inventories := getInventoriesFromFile(filepath)
	log.Printf("There are %d elves with inventories", len(inventories))
	return getTop3Totals(inventories)
}

// print top inventory calorie total
func solvePuzzle1(top3 [3]int) {
	maxCals := top3[0]
	fmt.Printf("Most calories: %d\n", maxCals)
}

// print sum of top 3 inventory inventory calorie totals
func solvePuzzle2(top3 [3]int) {
	totalCals := top3[0] + top3[1] + top3[2]
	fmt.Printf("Sum of Top 3: %d\n", totalCals)
}

func getInventoriesFromFile(filepath string) []string {
	normalizedInputBody := getFileAsString(filepath)
	return strings.Split(normalizedInputBody, "\n\n")
}

// get top 3 calorie totals from a set of inventories
func getTop3Totals(inventories []string) [3]int {
	var top3 [3]int

	for elfIdx, inventory := range inventories {
		currCals := getTotalCalories(inventory, elfIdx)
		if currCals > top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = currCals
		} else if currCals > top3[1] {
			top3[2] = top3[1]
			top3[1] = currCals
		} else if currCals > top3[2] {
			top3[2] = currCals
		}
	}
	return top3
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
