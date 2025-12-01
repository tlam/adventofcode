package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func main() {
	fp, _ := os.Open("input.txt")

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	minDial := 0
	maxDial := 99
	startDial := 50
	zeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			break
		}

		direction := line[:1]
		rotations, _ := strconv.Atoi(line[1:])

		if direction == "L" {
			startDial -= rotations
		} else {
			startDial += rotations
		}

		fmt.Printf("Start dial: %d\n", startDial)
		cycleCount := abs(startDial) / 100

		if cycleCount > 0 {
			totalCycleValues := cycleCount * 100

			if startDial > maxDial {
				startDial -= totalCycleValues
			} else if startDial < 0 {
				startDial += totalCycleValues
			}
		}

		if startDial < minDial {
			startDial = maxDial - abs(startDial) + 1
		} else if startDial > maxDial {
			startDial -= maxDial - 1
		}

		fmt.Printf("Turned %s %d: %d\n", direction, rotations, startDial)

		if startDial == 0 {
			zeroCount += 1
		}
	}

	fmt.Println(zeroCount)
}
