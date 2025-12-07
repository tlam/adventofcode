package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildWorksheet(filename string) [][]string {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var worksheet [][]string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		fmt.Println(line)

		worksheet = append(worksheet, strings.Split(line, ""))

	}

	return worksheet
}

// Parse the file as a string from top right to bottom right. Collect the number as a string down a column.
// Once an ops is found, stop and do the calculation.
func GrandTotal(filename string) int {
	total := 0

	worksheet := buildWorksheet(filename)

	fmt.Println(worksheet)

	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}

	column := []int{}

	for j := len(worksheet[0]) - 1; j >= 0; j-- {
		columnValue := ""
		for i := 0; i < len(worksheet); i++ {
			value := worksheet[i][j]
			if value == "+" || value == "*" {
				intValue, _ := strconv.Atoi(strings.TrimSpace(columnValue))
				if intValue > 0 {
					column = append(column, intValue)
				}
				fmt.Printf("Op %s for %v\n", value, column)

				columnTotal := 0

				if value == "*" {
					columnTotal = 1
				}

				for _, v := range column {
					columnTotal = ops[value](columnTotal, v)
				}

				total += columnTotal

				column = []int{}
				columnValue = ""
			} else {
				columnValue += value
			}
		}
		intValue, _ := strconv.Atoi(strings.TrimSpace(columnValue))
		if intValue > 0 {
			column = append(column, intValue)
		}
	}

	return total
}

func main() {
	output := GrandTotal("fixtures/input.txt")
	fmt.Println(output)
}
