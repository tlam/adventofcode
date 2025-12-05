package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fresh struct {
	lower int
	upper int
}

type Pantry []Fresh

// Build the slice of fresh ingredient ranges and available ingredients. Return pantry and slice of available ingredients
func buildPantry(filename string) (Pantry, []int) {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var pantry Pantry
	ingredients := []int{}

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
		} else {
			v, _ := strconv.Atoi(line)
			ingredients = append(ingredients, v)
		}

	}

	return pantry, ingredients
}

func FreshCount(filename string) int {
	count := 0

	pantry, ingredients := buildPantry(filename)

	for _, v := range ingredients {
		for _, fresh := range pantry {
			if v >= fresh.lower && v <= fresh.upper {
				count++
				break
			}

		}
	}

	return count
}

func main() {
	output := FreshCount("fixtures/input.txt")
	fmt.Println(output)
}
