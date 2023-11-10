package main

import (
	"testing"
)

const SINGLE_ITEM_FILE = "./tests/single-item-inv.txt"
const MULTI_ITEM_FILE = "./tests/multi-item-inv.txt"
const SINGLE_AND_MULTI_FILE = "./tests/single-and-multi-item-inv.txt"
const LESS_THAN_3_INVS_FILE = "./tests/less-than-3-invs.txt"

// region puzzle1
func Test_singleItemInvsPuzzle1(t *testing.T) {
	want := 55
	testResult := getTop3FromFile(SINGLE_ITEM_FILE)

	if want != testResult[0] {
		t.Fatalf("singleItemInvsPuzzle1.testResult = %d, want %d", testResult, want)
	}
}

func Test_multiItemInvsPuzzle1(t *testing.T) {
	want := 300
	testResult := getTop3FromFile(MULTI_ITEM_FILE)

	if want != testResult[0] {
		t.Fatalf("multiItemInvsPuzzle1.testResult = %d, want %d", testResult, want)
	}
}

func Test_singleAndMultiItemInvsPuzzle1(t *testing.T) {
	want := 1000
	testResult := getTop3FromFile(SINGLE_AND_MULTI_FILE)

	if want != testResult[0] {
		t.Fatalf("singleAndMultiItemInvsPuzzle1.testResult = %d, want %d", testResult, want)
	}
}

//endregion puzzle1

// region puzzle2
func Test_singleItemInvsPuzzle2(t *testing.T) {
	want := 55 + 22 + 12
	testResult := getTop3FromFile(SINGLE_ITEM_FILE)

	if want != testResult[0]+testResult[1]+testResult[2] {
		t.Fatalf("singleItemInvsPuzzle2.testResult = %d, want %d", testResult, want)
	}
}

func Test_multiItemInvsPuzzle2(t *testing.T) {
	want := 300 + 299 + 298
	testResult := getTop3FromFile(MULTI_ITEM_FILE)

	if want != testResult[0]+testResult[1]+testResult[2] {
		t.Fatalf("multiItemInvsPuzzle2.testResult = %d, want %d", testResult, want)
	}
}

func Test_singleAndMultiItemInvsPuzzle2(t *testing.T) {
	want := 1000 + 500 + 300
	testResult := getTop3FromFile(SINGLE_AND_MULTI_FILE)

	if want != testResult[0]+testResult[1]+testResult[2] {
		t.Fatalf("singleAndMultiItemInvsPuzzle2.testResult = %d, want %d", testResult, want)
	}
}

func Test_lessThan3InvsPuzzle2(t *testing.T) {
	want := 111 + 222
	testResult := getTop3FromFile(LESS_THAN_3_INVS_FILE)

	if want != testResult[0]+testResult[1]+testResult[2] {
		t.Fatalf("lessThan3InvsPuzzle2.testResult = %d, want %d", testResult, want)
	}
}

//endregion puzzle2
