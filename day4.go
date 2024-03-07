package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	cards := map[int]int{}
	reg := regexp.MustCompile("^Card +(\\d+)$")
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ":")
		cardNumber, _ := strconv.Atoi(reg.FindStringSubmatch(data[0])[1])
		cardNumbers := strings.Split(data[1], "|")
		winnings := prepareNumbers(cardNumbers[0], numbersRegex)
		numbers := prepareNumbers(cardNumbers[1], numbersRegex)
		nbWinnings := 0
		for _, number := range numbers {
			if inArray(number, winnings) {
				nbWinnings++
			}
		}
		v, _ := cards[cardNumber]
		cards[cardNumber] = v + 1
		for i := cardNumber + 1; i <= cardNumber+nbWinnings; i++ {
			fv, _ := cards[i]
			cards[i] = fv + cards[cardNumber]
		}
	}

	for k, v := range cards {
		fmt.Println(k, v)
		total += v
	}

	return total
}
