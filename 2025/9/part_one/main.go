package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func getRedTiles(filename string) []Coordinate {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var redTiles []Coordinate

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		coordinate := strings.Split(line, ",")

		x, _ := strconv.Atoi(coordinate[0])
		y, _ := strconv.Atoi(coordinate[1])

		redTiles = append(redTiles, Coordinate{x, y})

	}

	return redTiles
}

func absoluteValue(value int) int {
	if value < 0 {
		return value * -1
	}

	return value
}

// Given two opposite coordinates, calculate the width and height of the rectangle
func LargestRectangle(filename string) int {
	largestArea := 0

	redTiles := getRedTiles(filename)

	fmt.Println(redTiles)

	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			width := absoluteValue(redTiles[i].x-redTiles[j].x) + 1
			height := absoluteValue(redTiles[i].y-redTiles[j].y) + 1

			area := width * height

			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

func main() {
	output := LargestRectangle("fixtures/input.txt")
	fmt.Println(output)
}
