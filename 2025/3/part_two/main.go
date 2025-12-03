package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MaxBatteries = 12

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

func maxValue(subset []int) (int, int) {
	maxSoFar := -1
	index := -1

	for i, v := range subset {
		if v > maxSoFar {
			index = i
			maxSoFar = v
		}
	}

	return index, maxSoFar
}

// Find the largest number in the  first n - MaxBatteries items.
// Once found, advance the index and keep searching in the remainder of the array.
// The search is complete if there's only one item left or the remainder is empty.
func LargestJoltage(bank []int) int {
	batteries := []int{}
	count := 0
	indexSoFar := 0
	subset := bank[:len(bank)-MaxBatteries+1]

	fmt.Printf("Bank: %v\n", bank)

	for count < MaxBatteries {
		fmt.Printf("Searching for next number in %v\n", subset)

		if len(subset) == 1 {
			batteries = append(batteries, subset[0])
			break
		}

		maxIndex := -1
		maxSoFar := -1

		maxIndex, maxSoFar = maxValue(subset)
		fmt.Printf("Max so far: %d, max index: %d\n", maxSoFar, maxIndex)
		batteries = append(batteries, maxSoFar)

		fmt.Printf("Batteries: %v\n", batteries)

		if count == 0 {
			indexSoFar += maxIndex
		} else {
			indexSoFar += maxIndex + 1
		}

		subset = bank[indexSoFar+1:]

		if len(batteries)+len(subset) == MaxBatteries {
			batteries = append(batteries, subset...)
			break
		}

		count++

		startIndex := indexSoFar

		offset := 1 + len(bank) - MaxBatteries + count

		if offset > len(bank) {
			subset = bank[len(bank)-1:]
		} else if offset > startIndex {
			subset = bank[startIndex+1 : offset]
		}

	}

	sum := ""

	for i := 0; i < len(batteries); i++ {
		sum += strconv.Itoa(batteries[i])
	}

	sumInt := 0
	sumInt, _ = strconv.Atoi(sum)

	return sumInt

}

func TotalJoltage(filename string) int {
	sum := 0

	banks := getBanks(filename)

	for _, bank := range banks {
		joltage := LargestJoltage(bank)
		fmt.Printf("Got joltage %d\n", joltage)

		sum += joltage
	}

	return sum
}

func main() {
	output := TotalJoltage("fixtures/input.txt")
	fmt.Println(output)
}
