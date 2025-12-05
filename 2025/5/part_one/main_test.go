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
	count := FreshCount("fixtures/sample.txt")
	AssertEqual(t, count, 3)
}
