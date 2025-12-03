package main

import (
	"testing"
)

func AssertEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("Expected %d, got %d", expected, value)
	}
}

func TestZeroCountR50(t *testing.T) {
	count := ZeroCount("fixtures/r50.txt")

	AssertEqual(t, count, 1)
}

func TestZeroCountR150(t *testing.T) {
	count := ZeroCount("fixtures/r150.txt")

	AssertEqual(t, count, 2)
}

func TestZeroCountL150(t *testing.T) {
	count := ZeroCount("fixtures/l150.txt")

	AssertEqual(t, count, 2)
}

func TestZeroCountL75(t *testing.T) {
	count := ZeroCount("fixtures/l75.txt")

	AssertEqual(t, count, 1)
}

func TestZeroCountR75(t *testing.T) {
	count := ZeroCount("fixtures/r75.txt")

	AssertEqual(t, count, 1)
}

func TestZeroCountSample(t *testing.T) {
	count := ZeroCount("sample.txt")

	AssertEqual(t, count, 6)
}

func TestZeroCountInput(t *testing.T) {
	count := ZeroCount("input.txt")

	AssertEqual(t, count, 5831)
}
