package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day2(input string) int {
	total := 0
	reg := regexp.MustCompile("(\\d+) (\\w+)")

	fromInput(input, func(line string) {
		// get game and data
		gameParts := strings.Split(line, ":")
		gameNumber, err := strconv.Atoi(strings.Replace(gameParts[0], "Game ", "", 1))
		if err != nil {
			panic(err)
		}
		sets := strings.Split(gameParts[1], ";")
		impossible := false
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				vals := reg.FindStringSubmatch(cube)
				n, _ := strconv.Atoi(vals[1])
				switch vals[2] {
				case "red":
					if n > 12 {
						impossible = true
					}
				case "blue":
					if n > 14 {
						impossible = true
					}
				case "green":
					if n > 13 {
						impossible = true
					}
				}
				if impossible {
					break
				}
			}
			if impossible {
				break
			}
		}
		if !impossible {
			total += gameNumber
		}
	})

	return total
}

func day2_2(input string) int {
	total := 0
	reg := regexp.MustCompile("(\\d+) (\\w+)")

	fromInput(input, func(line string) {
		// get game and data
		gameParts := strings.Split(line, ":")
		_, err := strconv.Atoi(strings.Replace(gameParts[0], "Game ", "", 1))
		if err != nil {
			panic(err)
		}
		sets := strings.Split(gameParts[1], ";")
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				vals := reg.FindStringSubmatch(cube)
				n, _ := strconv.Atoi(vals[1])
				switch vals[2] {
				case "red":
					if n > maxRed {
						maxRed = n
					}
				case "blue":
					if n > maxBlue {
						maxBlue = n
					}
				case "green":
					if n > maxGreen {
						maxGreen = n
					}
				}
			}
		}
		gameResult := maxRed * maxBlue * maxGreen
		total += gameResult
	})

	return total
}
