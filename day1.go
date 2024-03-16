package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day1() int {
	total := 0
	reg := regexp.MustCompile(`(\d)`)
	fromInput("inputs/day1.txt", func(line string) {
		digits := reg.FindAllStringSubmatch(line, -1)
		first := digits[0][0]
		last := digits[len(digits)-1][0]
		number, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			panic(err)
		}
		total += number
	})
	return total
}

var numberMapping = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var neededs = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func day1_2() int {
	total := 0
	fromInput("inputs/day1.txt", func(line string) {
		first := ""
		lowerIndex := len(line)
		greaterIndex := 0
		last := ""
		for _, n := range neededs {
			lastV := strings.LastIndex(line, n)
			firstV := strings.Index(line, n)
			if firstV != -1 && firstV <= lowerIndex {
				lowerIndex = firstV
				first = n
			}
			if lastV != -1 && lastV >= greaterIndex {
				greaterIndex = lastV
				last = n
			}
		}
		number, err := strconv.Atoi(fmt.Sprintf("%d%d", getDigit(first), getDigit(last)))
		if err != nil {
			panic(err)
		}

		total += number
	})
	return total
}

func getDigit(v string) int {
	iv, err := strconv.Atoi(v)
	if err != nil {
		iv = numberMapping[v]
	}

	return iv
}
