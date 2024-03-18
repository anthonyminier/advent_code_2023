package main

import (
	"strings"
)

func day6(input string) int {
	times := []int{}
	distances := []int{}
	// format the data
	fromInput(input, func(line string) {
		parts := strings.Split(line, ":")
		values := parseNumberSuits(parts[1])
		if strings.Contains(parts[0], "Time") {
			times = values
		}
		if strings.Contains(parts[0], "Distance") {
			distances = values
		}
	})
	total := 1

	for i, t := range times {
		nbWaysToWin := 0
		for nbMilli := range t {
			// distance = nb millisecond to run * nb milliseconds holds
			dist := (t - nbMilli) * nbMilli
			if dist > distances[i] {
				nbWaysToWin++
			}
		}
		if nbWaysToWin != 0 {
			total *= nbWaysToWin
		}
	}
	return total
}

func day6_2(input string) int {
	return day6(input)
}
