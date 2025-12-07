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
	AssertEqual(t, count, 14)
}

func TestRightOverlap(t *testing.T) {
	count := FreshCount("fixtures/right_overlap.txt")
	AssertEqual(t, count, 6)
}

func TestSubset(t *testing.T) {
	count := FreshCount("fixtures/subset.txt")
	AssertEqual(t, count, 5)
}

func TestLeftEdgeOverlp(t *testing.T) {
	count := FreshCount("fixtures/left_edge_overlap.txt")
	AssertEqual(t, count, 8)
}

func TestRightEdgeOverlp(t *testing.T) {
	count := FreshCount("fixtures/right_edge_overlap.txt")
	AssertEqual(t, count, 5)
}

func TestEdgesOverlap(t *testing.T) {
	count := FreshCount("fixtures/edges.txt")
	AssertEqual(t, count, 12)
}

func TestApart(t *testing.T) {
	count := FreshCount("fixtures/apart.txt")
	AssertEqual(t, count, 6)
}

func TestMultipleSubsets(t *testing.T) {
	count := FreshCount("fixtures/multiple_subsets.txt")
	AssertEqual(t, count, 7)
}

func TestRemoveDuplicates(t *testing.T) {
	count := FreshCount("fixtures/duplicates.txt")
	AssertEqual(t, count, 6)
}
