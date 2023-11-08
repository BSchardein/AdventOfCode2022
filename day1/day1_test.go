package main

import (
	"testing"
)

const SINGLE_ITEM_FILE = "./tests/single-item-inv.txt"
const MULTI_ITEM_FILE = "./tests/multi-item-inv.txt"
const SINGLE_AND_MULTI_FILE = "./tests/single-and-multi-item-inv.txt"

func Test_singleItemInvs(t *testing.T) {
	want := 55
	testResult := getMaxCalsFromFile(SINGLE_ITEM_FILE)

	if want != testResult {
		t.Fatalf("singleItemInvs.testResult = %d, want %d", testResult, want)
	}
}

func Test_multiItemInvs(t *testing.T) {
	want := 300
	testResult := getMaxCalsFromFile(MULTI_ITEM_FILE)

	if want != testResult {
		t.Fatalf("multiItemInvs.testResult = %d, want %d", testResult, want)
	}
}

func Test_singleAndMultiItemInvs(t *testing.T) {
	want := 1000
	testResult := getMaxCalsFromFile(SINGLE_AND_MULTI_FILE)

	if want != testResult {
		t.Fatalf("singleAndMultiItemInvs.testResult = %d, want %d", testResult, want)
	}
}
