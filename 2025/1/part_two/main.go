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

func ZeroCount(filename string) int {
	fp, _ := os.Open(filename)

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
		initialStartDial := startDial
		cycleCount := rotations / 100

		if cycleCount > 0 {
			rotations = rotations - (cycleCount * 100)

			fmt.Printf("Remainder is %d\n", rotations)

		}

		if direction == "L" {
			startDial -= rotations
			// After going backward with the rotations, the value is negative integer. Zero has been crossed.
			if initialStartDial != 0 && startDial < minDial {
				zeroCount++
			}
		} else { // R - Forward
			startDial += rotations
			// After going forward with the rotations, the value is a positive integer. Zero has been crossed.
			if initialStartDial != 0 && startDial > 100 {
				zeroCount++
			}
		}

		if startDial < minDial {
			startDial = maxDial - abs(startDial) + 1
		} else if startDial == 100 {
			startDial = 0
		} else if startDial > maxDial {
			startDial = startDial - maxDial - 1
		}

		if startDial == 0 {
			zeroCount++
		}

		fmt.Printf("Turned %s %d, cycles: %d: %d, zero count so far: %d\n\n", direction, rotations, cycleCount, startDial, zeroCount)

		// Each full cycle crosses zero.
		zeroCount += cycleCount

	}

	return zeroCount
}

func main() {
	filename := "input.txt"
	output := ZeroCount(filename)
	fmt.Println(output)
}
