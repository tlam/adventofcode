package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildGrid(filename string) (int, [][]string) {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var grid [][]string

	lineNumber := 0
	manifoldLocation := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if lineNumber == 0 {
			manifoldLocation = strings.Index(line, "S")
		}

		grid = append(grid, strings.Split(line, ""))
		lineNumber++
	}

	return manifoldLocation, grid
}

// When a splitter is found, mark the left and right as `|` if they are `.`.
// In the recursive call, if the splitter left or right already has a `|`, do not go down that path
// An earlier function call has previously followed that path.
func Split(i int, j int, grid [][]string) int {
	if i >= len(grid) || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] == "^" {
		if grid[i][j-1] == "." && grid[i][j+1] == "." {
			grid[i][j-1] = "|"
			grid[i][j+1] = "|"
			return 1 + Split(i+2, j-1, grid) + Split(i+2, j+1, grid)
		} else if grid[i][j-1] == "." && grid[i][j+1] == "|" {
			grid[i][j-1] = "|"
			return 1 + Split(i+2, j-1, grid)
		} else if grid[i][j-1] == "|" && grid[i][j+1] == "." {
			grid[i][j+1] = "|"
			return 1 + Split(i+2, j+1, grid)
		} else {
			return 0
		}
	}

	return Split(i+2, j, grid)
}

func SplitCount(filename string) int {
	manifoldLocation, grid := buildGrid(filename)

	count := Split(2, manifoldLocation, grid)

	return count
}

func main() {
	output := SplitCount("fixtures/input.txt")
	fmt.Println(output)
}
