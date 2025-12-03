package main

import (
	"testing"
)

func AssertEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("Expected %d, got %d", expected, value)
	}
}

func TestSampleInvalidIds(t *testing.T) {
	sum := InvalidIds("fixtures/sample.txt")

	AssertEqual(t, sum, 4174379265)
}

func TestCompareEleven(t *testing.T) {
	count := Compare("11", 0, "", "", 0)

	AssertEqual(t, count, 2)
}

func TestCompare110(t *testing.T) {
	count := Compare("110", 0, "", "", 0)

	AssertEqual(t, count, 1)
}

func TestCompare1111111(t *testing.T) {
	count := Compare("1111111", 0, "", "", 0)

	AssertEqual(t, count, 7)
}

func TestCompare12341234(t *testing.T) {
	count := Compare("12341234", 0, "", "", 0)

	AssertEqual(t, count, 2)
}

func TestCompare123123123(t *testing.T) {
	count := Compare("123123123", 0, "", "", 0)

	AssertEqual(t, count, 3)
}

func TestCompare1188511885(t *testing.T) {
	count := Compare("1188511885", 0, "", "", 0)

	AssertEqual(t, count, 2)
}

func TestCompare565656(t *testing.T) {
	count := Compare("565656", 0, "", "", 0)

	AssertEqual(t, count, 3)
}

func TestCompare74740(t *testing.T) {
	count := Compare("74740", 0, "", "", 0)

	AssertEqual(t, count, 1)
}
