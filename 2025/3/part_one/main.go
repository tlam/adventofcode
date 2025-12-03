package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Convert the string of numbers into an int slice
func getBanks(filename string) [][]int {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	banks := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			break
		}

		bank := []int{}

		for i := 0; i < len(line); i++ {
			bank = append(bank, int(line[i]-'0'))
		}

		banks = append(banks, bank)
	}

	return banks
}

// Set first and second based on the first two numbers.
// If a larger first is encountered later, update first and set second to the next one if possible
// Continue checking if second has changed as each number is being iterated
func LargestJoltage(bank []int) int {
	first := bank[0]
	second := -1

	bankLength := len(bank)
	i := 0

	for i < bankLength {
		j := i + 1

		if j < bankLength {
			if bank[j] > second {
				second = bank[j]
			}
		} else {
			joltage, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
			return joltage
		}

		fmt.Printf("First %d, second %d\n", first, second)

		if second > first {
			if j == bankLength-1 {

			} else {
				first = second
				second = bank[j+1]
			}
		}

		i++

	}

	return 0

}

func TotalJoltage(filename string) int {
	sum := 0

	banks := getBanks(filename)

	for _, bank := range banks {
		joltage := LargestJoltage(bank)
		fmt.Println(bank)
		fmt.Printf("Got joltage %d\n", joltage)

		sum += joltage
	}

	return sum
}

func main() {
	output := TotalJoltage("fixtures/input.txt")
	fmt.Println(output)
}
