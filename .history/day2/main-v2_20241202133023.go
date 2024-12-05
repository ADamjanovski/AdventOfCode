package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// checkReportConsistency checks if a report is consistently increasing or decreasing
func checkReportConsistency(levels []int) bool {
	if len(levels) <= 1 {
		return true
	}

	// Try increasing pattern
	isIncreasing := true
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < 1 || diff > 3 {
			isIncreasing = false
			break
		}
	}
	if isIncreasing {
		return true
	}

	// Try decreasing pattern
	isDecreasing := true
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		if diff < 1 || diff > 3 {
			isDecreasing = false
			break
		}
	}
	return isDecreasing
}

// isSafeWithProblemDampener checks if a report is safe, optionally allowing one level removal
func isSafeWithProblemDampener(levels []int) bool {
	// Check if report is safe without removing any level
	if checkReportConsistency(levels) {
		return true
	}

	// Try removing each level
	for i := 0; i < len(levels); i++ {
		// Create a new slice without the i-th element
		modifiedLevels := make([]int, 0, len(levels)-1)
		modifiedLevels = append(modifiedLevels, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)

		// Check if modified report is safe
		if checkReportConsistency(modifiedLevels) {
			return true
		}
	}

	return false
}

func main() {
	f, err := os.Open("data-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	safeLevels := 0

	for scanner.Scan() {
		// Parse the line into integers
		parts := strings.Fields(scanner.Text())
		levels := make([]int, len(parts))
		
		for i, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error parsing level:", err)
				break
			}
			levels[i] = level
		}

		// Check if the report is safe with Problem Dampener
		if isSafeWithProblemDampener(levels) {
			safeLevels++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Safe Levels:", safeLevels)
}