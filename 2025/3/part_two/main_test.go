package main

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("Expected %d, got %d", expected, value)
	}
}

func TestSample(t *testing.T) {
	joltage := TotalJoltage("fixtures/sample.txt")
	AssertEqual(t, joltage, 3121910778619)
}

func TestLargestJoltage987654321111111(t *testing.T) {
	bank := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 987654321111)
}

func TestLargestJoltage811111111111119(t *testing.T) {
	bank := []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 811111111119)
}

func TestLargestJoltage234234234234278(t *testing.T) {
	bank := []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 434234234278)
}

func TestLargestJoltage818181911112111(t *testing.T) {
	bank := []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}
	joltage := LargestJoltage(bank)

	AssertEqual(t, joltage, 888911112111)
}

func TestLargestJoltageSample(t *testing.T) {
	bank := []int{5, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 5, 2, 3, 4, 2, 1, 2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 1, 2, 3, 3, 6, 1, 2, 2, 2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1, 1, 3, 2, 2, 2, 2, 2, 1, 3, 2, 3, 1, 2, 4, 2, 2, 3, 2, 2, 2, 2, 3, 3, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1}
	joltage := LargestJoltage(bank)

	fmt.Println(joltage)

	AssertEqual(t, joltage, 643332222222)
}
