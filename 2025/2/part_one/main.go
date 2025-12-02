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

func InvalidIds(filename string) int {
	sum := 0

	ranges := getRanges(filename)

	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")
		firstId := ids[0]
		secondId := ids[1]

		// Skip all ranges where the ids have the same length and are of odd length.
		if len(firstId) == len(secondId) && len(firstId)%2 != 0 {
			fmt.Printf("Skipping %s\n", idRange)
			continue
		}

		minId, _ := strconv.Atoi(firstId)
		maxId, _ := strconv.Atoi(secondId)

		for i := range maxId - minId + 1 {
			intValue := i + minId
			value := strconv.Itoa(intValue)
			median := len(value) / 2

			firstHalf := value[0:median]
			secondHalf := value[median:]

			fmt.Printf("Value: %s, %s %s\n", value, firstHalf, secondHalf)

			if firstHalf == secondHalf {
				fmt.Printf("%s has invalid ID %s\n", idRange, value)
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
