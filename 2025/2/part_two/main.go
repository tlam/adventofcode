// For sequence of digits that are repeated twice, the length of the id must be even.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRanges(filename string) []string {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	ranges := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			break
		}

		ranges = strings.Split(line, ",")
	}

	return ranges
}

func Compare(originalString string, index int, pattern string, substring string, count int) int {
	if index == 0 {
		pattern = string(originalString[:1])
		substring = originalString[1:]
		count++
	}

	index++

	patternLength := len(pattern)

	if patternLength > len(substring) {
		if substring == "" {
			return count
		} else {
			return 1
		}
	}

	//fmt.Printf("Pattern %s, compare substring %s, substring %s: %d\n", pattern, substring[:patternLength], substring, count)

	if pattern == substring[:patternLength] {
		return Compare(originalString, index, pattern, substring[patternLength:], count+1)
	} else {
		pattern = originalString[:index]
		return Compare(originalString, index, pattern, originalString[index:], 1)
	}
}

func InvalidIds(filename string) int {
	sum := 0

	ranges := getRanges(filename)

	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")
		firstId := ids[0]
		secondId := ids[1]

		minId, _ := strconv.Atoi(firstId)
		maxId, _ := strconv.Atoi(secondId)

		for i := range maxId - minId + 1 {
			intValue := i + minId
			strValue := strconv.Itoa(intValue)

			patternCount := Compare(strValue, 0, "", "", 0)

			if patternCount >= 2 {
				fmt.Printf("Range %s, Invalid ID: %d\n", idRange, intValue)

				sum += intValue
			}
		}

	}

	return sum
}

func main() {
	output := InvalidIds("fixtures/input.txt")
	fmt.Println(output)
}
