package main

import (
	"testing"
)

func TestZeroCountR50(t *testing.T) {
	count := ZeroCount("fixtures/r50.txt")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestZeroCountR150(t *testing.T) {
	count := ZeroCount("fixtures/r150.txt")
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestZeroCountL150(t *testing.T) {
	count := ZeroCount("fixtures/l150.txt")
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestZeroCountL75(t *testing.T) {
	count := ZeroCount("fixtures/l75.txt")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestZeroCountR75(t *testing.T) {
	count := ZeroCount("fixtures/r75.txt")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestZeroCountSample(t *testing.T) {
	count := ZeroCount("sample.txt")
	if count != 6 {
		t.Errorf("Expected 6, got %d", count)
	}
}

func TestZeroCountInput(t *testing.T) {
	count := ZeroCount("input.txt")
	if count != 5831 {
		t.Errorf("Expected 5831, got %d", count)
	}
}
