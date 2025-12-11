package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildDevices(filename string) map[string][]string {
	fp, _ := os.Open(filename)

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	devices := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		splitLine := strings.Split(line, ":")
		outputs := strings.Split(strings.TrimSpace(splitLine[1]), " ")

		devices[splitLine[0]] = outputs
	}

	return devices
}

// Recursively go through each possible path and count if the final output is out
func escape(devices map[string][]string, k string, count *int) {
	outputs, exists := devices[k]

	if !exists {
		return
	}

	if outputs[0] == "out" {
		*count++
	} else {
		for _, output := range outputs {
			escape(devices, output, count)
		}
	}
}

func PathCount(filename string) int {
	count := 0
	devices := buildDevices(filename)
	escape(devices, "you", &count)
	return count
}

func main() {
	count := PathCount("fixtures/input.txt")
	fmt.Println(count)
}
