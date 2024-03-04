package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func day4(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	numbersRegex := regexp.MustCompile("(\\d+)")
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ":")
		cardNumbers := strings.Split(data[1], "|")
		winnings := prepareNumbers(cardNumbers[0], numbersRegex)
		numbers := prepareNumbers(cardNumbers[1], numbersRegex)
		subTotal := 0
		for _, n := range numbers {
			if inArray(n, winnings) {
				if subTotal == 0 {
					subTotal = 1
				} else {
					subTotal *= 2
				}
			}
		}

		total += subTotal
	}

	return total
}

func inArray[T comparable](needed T, elements []T) bool {
	for _, n := range elements {
		if n == needed {
			return true
		}
	}
	return false
}

func prepareNumbers(numbers string, reg *regexp.Regexp) []string {
	matches := reg.FindAllStringSubmatch(numbers, -1)
	results := []string{}
	for _, n := range matches {
		results = append(results, n[0])
	}

	return results
}

func day4_2(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	numbersRegex := regexp.MustCompile("(\\d+)")
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ":")
		cardNumbers := strings.Split(data[1], "|")
		winnings := prepareNumbers(cardNumbers[0], numbersRegex)
		numbers := prepareNumbers(cardNumbers[1], numbersRegex)
		subTotal := 0
		for _, n := range numbers {
			if inArray(n, winnings) {
				if subTotal == 0 {
					subTotal = 1
				} else {
					subTotal *= 2
				}
			}
		}

		total += subTotal
	}

	return total
}
