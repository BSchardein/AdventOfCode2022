package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const FILEPATH = "./input.txt"

var playerChoiceMap = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var winCaseMap = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var choiceValueMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var LOSSVALUE = 0
var DRAWVALUE = 3
var WINVALUE = 6

//A = X = Rock 			= 1
//B = Y = Paper 		= 2
//C = Z = Scissors 	= 3

func main() {
	puzzle1Solution := solvePuzzle1(FILEPATH)
	fmt.Printf("Total Score from Guide: %d", puzzle1Solution)
}

func solvePuzzle1(filePath string) int {
	totalScore := 0
	battles := strings.Split(getFileAsString(filePath), "\n")
	for battleIdx, battle := range battles {
		if len(battle) > 0 {
			fmt.Printf("Battle %d: %s\n", battleIdx, battle)
			moves := strings.Split(battle, " ")
			totalScore += getBattleOutcome(moves[0], playerChoiceMap[moves[1]])
		}
	}
	return totalScore
}

func getBattleOutcome(opponentMove string, playerMove string) int {
	outcome := LOSSVALUE
	if playerMove == opponentMove {
		outcome = DRAWVALUE
	} else if winCaseMap[playerMove] == opponentMove {
		outcome = WINVALUE
	}
	return outcome + choiceValueMap[playerMove]
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
