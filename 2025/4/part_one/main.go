package main

import (
	"bufio"
	"fmt"
	"os"
)

const PaperRoll string = "@"
const RollThreshold int = 4

// Convert the string of numbers into an int slice
func get2dArray(filename string) [][]string {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var rolls [][]string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			break
		}

		var row []string

		for _, r := range line {
			row = append(row, string(r))
		}

		rolls = append(rolls, row)
	}

	return rolls
}

func indexValid(i int, j int, rowCount int, columnCount int) bool {
	return i >= 0 && j >= 0 && i < rowCount && j < columnCount
}

// Given a roll at a position, check other rolls relative to it clockwise. Start from N, NE, E, ...
func rollAccessible(rolls [][]string, i int, j int) bool {
	roll := rolls[i][j]

	if roll == "." {
		return false
	}

	rowCount := len(rolls)
	columnCount := len(rolls[0])
	rollCount := 0

	// North check
	ni := i - 1

	if indexValid(ni, j, rowCount, columnCount) && rolls[ni][j] == PaperRoll {
		rollCount++
	}

	// North East check
	nei := i - 1
	nej := j + 1

	if indexValid(nei, nej, rowCount, columnCount) && rolls[nei][nej] == PaperRoll {
		rollCount++
	}

	// East check
	ej := j + 1

	if indexValid(i, ej, rowCount, columnCount) && rolls[i][ej] == PaperRoll {
		rollCount++
	}

	// South East check
	sei := i + 1
	sej := j + 1

	if indexValid(sei, sej, rowCount, columnCount) && rolls[sei][sej] == PaperRoll {
		rollCount++
	}

	// South check
	si := i + 1

	if indexValid(si, j, rowCount, columnCount) && rolls[si][j] == PaperRoll {
		rollCount++
	}

	// South West check
	swi := i + 1
	swj := j - 1

	if indexValid(swi, swj, rowCount, columnCount) && rolls[swi][swj] == PaperRoll {
		rollCount++
	}

	// West check
	wj := j - 1

	if indexValid(i, wj, rowCount, columnCount) && rolls[i][wj] == PaperRoll {
		rollCount++
	}

	// North West check
	nwi := i - 1
	nwj := j - 1

	if indexValid(nwi, nwj, rowCount, columnCount) && rolls[nwi][nwj] == PaperRoll {
		rollCount++
	}

	return rollCount < RollThreshold
}

func RollCount(filename string) int {
	sum := 0

	rolls := get2dArray(filename)

	for i := 0; i < len(rolls); i++ {
		for j := 0; j < len(rolls[i]); j++ {
			if rollAccessible(rolls, i, j) {
				sum += 1
			}
		}

	}

	return sum
}

func main() {
	output := RollCount("fixtures/input.txt")
	fmt.Println(output)
}
