package main

import (
	"testing"
)

func TestSampleInvalidIds(t *testing.T) {
	sum := InvalidIds("fixtures/sample.txt")
	expectedSum := 4174379265

	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestCompareEleven(t *testing.T) {
	count := Compare("11", 0, "", "", 0)
	expectedCount := 2

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare110(t *testing.T) {
	count := Compare("110", 0, "", "", 0)
	expectedCount := 1

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare1111111(t *testing.T) {
	count := Compare("1111111", 0, "", "", 0)
	expectedCount := 7

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare12341234(t *testing.T) {
	count := Compare("12341234", 0, "", "", 0)
	expectedCount := 2

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare123123123(t *testing.T) {
	count := Compare("123123123", 0, "", "", 0)
	expectedCount := 3

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare1188511885(t *testing.T) {
	count := Compare("1188511885", 0, "", "", 0)
	expectedCount := 2

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare565656(t *testing.T) {
	count := Compare("565656", 0, "", "", 0)
	expectedCount := 3

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}

func TestCompare74740(t *testing.T) {
	count := Compare("74740", 0, "", "", 0)
	expectedCount := 1

	if count != expectedCount {
		t.Errorf("Expected %d, got %d", expectedCount, count)
	}
}
