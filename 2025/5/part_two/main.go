package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Fresh struct {
	lower int
	upper int
}

type Pantry []Fresh

// Build the slice of fresh ingredient ranges.
func buildPantry(filename string) Pantry {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var pantry Pantry

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			values := strings.Split(line, "-")
			lower, _ := strconv.Atoi(values[0])
			upper, _ := strconv.Atoi(values[1])

			fresh := Fresh{lower, upper}
			pantry = append(pantry, fresh)
		}

	}

	return pantry
}

// Sort the ranges and merge the ones that are a subset of each other
func FreshCount(filename string) int {
	pantry := buildPantry(filename)

	total := 0

	sort.Slice(pantry, func(i, j int) bool {
		return pantry[i].lower < pantry[j].lower
	})

	current := pantry[0]
	i := 1

	for i < len(pantry) {
		next := pantry[i]

		if current.upper < next.lower {
			total += current.upper - current.lower + 1
			current = next
		} else { // subset found
			newSlice := Fresh{min(current.lower, next.lower), max(current.upper, next.upper)}
			current = newSlice
		}

		i++
	}
	total += current.upper - current.lower + 1

	return total
}

func main() {
	output := FreshCount("fixtures/input.txt")
	fmt.Println(output)
}
