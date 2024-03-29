package main

import (
	"strconv"
	"strings"
)

type data struct {
	source      int
	destination int
	nrange      int
}

func day5(input string) int {
	minimum := 1000000000000
	seedMap := [7][]data{}
	currentIndex := -1
	seeds := []int{}
	// format the data
	fromInput(input, func(line string) {
		if line == "" {
			return
		}
		if strings.Contains(line, "seeds:") {
			parse := strings.Split(line, " ")
			parts := parse[1:]
			for _, s := range parts {
				v, _ := strconv.Atoi(s)
				seeds = append(seeds, v)
			}
			return
		}
		if strings.Contains(line, "map:") {
			currentIndex++
			return
		}
		parts := strings.Split(line, " ")
		dest, _ := strconv.Atoi(parts[0])
		src, _ := strconv.Atoi(parts[1])
		nrange, _ := strconv.Atoi(parts[2])
		seedMap[currentIndex] = append(seedMap[currentIndex], data{source: src, destination: dest, nrange: nrange})
	})

	// compute path for each seed
	for _, seed := range seeds {
		value := seed
		for _, m := range seedMap {
			for _, d := range m {
				if value >= d.source && value < d.source+d.nrange {
					inc := d.destination - d.source
					value = value + inc
					break
				}
			}
		}
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

type seed struct {
	min int
	max int
}

// this solution (brute force one) take around 2 minutes to find the solution, not really optimized
// but i keep it and maybe later, i will try to find a better solution
func day5_2(input string) int {
	minimum := 1000000000000
	seedMap := [7][]data{}
	currentIndex := -1
	seeds := []seed{}
	fromInput(input, func(line string) {
		if line == "" {
			return
		}
		if strings.Contains(line, "seeds:") {
			parse := strings.Split(line, " ")
			parts := parse[1:]
			min := 0
			for _, s := range parts {
				v, _ := strconv.Atoi(s)
				if min == 0 {
					min = v
				} else {
					seeds = append(seeds, seed{min: min, max: v})
					min = 0
				}

			}
			return
		}
		if strings.Contains(line, "map:") {
			currentIndex++
			return
		}
		parts := strings.Split(line, " ")
		dest, _ := strconv.Atoi(parts[0])
		src, _ := strconv.Atoi(parts[1])
		nrange, _ := strconv.Atoi(parts[2])
		seedMap[currentIndex] = append(seedMap[currentIndex], data{source: src, destination: dest, nrange: nrange})
	})

	for _, seed := range seeds {
		for seedValue := seed.min; seedValue < seed.min+seed.max; seedValue++ {
			value := seedValue
			for _, m := range seedMap {
				for _, d := range m {
					if value >= d.source && value < d.source+d.nrange {
						inc := d.destination - d.source
						value = value + inc
						break
					}
				}
			}
			if value < minimum {
				minimum = value
			}
		}
	}

	return minimum
}
