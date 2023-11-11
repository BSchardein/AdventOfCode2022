package main

import (
	"fmt"
	"testing"
)

const BATTLE_FILE = "./tests/%s%s.txt"
const A = "A"
const B = "B"
const C = "C"
const WIN = "Win"
const LOSS = "Loss"
const DRAW = "Draw"
const MULTIBATTLE = "./tests/MultiBattle.txt"

// region puzzle1

func Test_MultiBattlePuzzle1(t *testing.T) {
	// rock loss, scissors draw, paper win
	round1 := 1 + 0
	round2 := 3 + 3
	round3 := 2 + 6
	want := round1 + round2 + round3
	result := solvePuzzle1(MULTIBATTLE)
	checkFailure(t, result, want)
}

func Test_ABattlesPuzzle1(t *testing.T) {
	playerMove := A
	checkWinLossDraw(t, playerMove)
}

func Test_BBattlesPuzzle1(t *testing.T) {
	playerMove := B
	checkWinLossDraw(t, playerMove)
}

func Test_CBattlesPuzzle1(t *testing.T) {
	playerMove := C
	checkWinLossDraw(t, playerMove)
}

func checkWinLossDraw(t *testing.T, move string) {
	moveValue := choiceValueMap[move]
	lossWant := moveValue + LOSSVALUE
	drawWant := moveValue + DRAWVALUE
	winWant := moveValue + WINVALUE
	checkMoveCase(t, move, LOSS, lossWant)
	checkMoveCase(t, move, DRAW, drawWant)
	checkMoveCase(t, move, WIN, winWant)
}

func checkMoveCase(t *testing.T, move string, outcome string, want int) {
	file := fmt.Sprintf(BATTLE_FILE, move, outcome)
	result := solvePuzzle1(file)
	checkFailure(t, result, want)
}

//endregion puzzle1

func checkFailure(t *testing.T, testResult int, want int) {
	if want != testResult {
		t.Fatalf("testResult = %d, want %d", testResult, want)
	}
}
