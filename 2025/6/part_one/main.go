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

		worksheet = append(worksheet, strings.Fields(line))

	}

	return worksheet
}

func GrandTotal(filename string) int {
	total := 0

	worksheet := buildWorksheet(filename)

	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}

	for i := 0; i < len(worksheet[0]); i++ {
		operator := worksheet[len(worksheet)-1][i]
		fmt.Printf("Operator %s\n", operator)

		columnTotal := 0

		if operator == "*" {
			columnTotal = 1
		}

		for j := 0; j < len(worksheet)-1; j++ {
			value, _ := strconv.Atoi(worksheet[j][i])

			columnTotal = ops[operator](columnTotal, value)

		}

		total += columnTotal

	}

	return total
}

func main() {
	output := GrandTotal("fixtures/input.txt")
	fmt.Println(output)
}
