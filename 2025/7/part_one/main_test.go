package main

import (
	"testing"
)

func AssertEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("Expected %d, got %d", expected, value)
	}
}

func TestSample(t *testing.T) {
	count := SplitCount("fixtures/sample.txt")
	AssertEqual(t, count, 21)
}

func TestSmall(t *testing.T) {
	count := SplitCount("fixtures/small.txt")
	AssertEqual(t, count, 6)
}
