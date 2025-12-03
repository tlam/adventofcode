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
	joltage := TotalJoltage("fixtures/sample.txt")
	AssertEqual(t, joltage, 357)
}

func TestLargestJoltage1234(t *testing.T) {
	bank := []int{1, 2, 3, 4}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 34)
}

func TestLargestJoltage4321(t *testing.T) {
	bank := []int{4, 3, 2, 1}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 43)
}

func TestLargestJoltage131516(t *testing.T) {
	bank := []int{1, 3, 1, 5, 1, 6}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 56)
}

func TestLargestJoltage818181911112111(t *testing.T) {
	bank := []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 92)
}
