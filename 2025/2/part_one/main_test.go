package main

import (
	"testing"
)

func TestInvalidIds(t *testing.T) {
	sum := InvalidIds("fixtures/sample.txt")
	if sum != 1227775554 {
		t.Errorf("Expected 1227775554, got %d", sum)
	}
}
