package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day1() int {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	reg := regexp.MustCompile("(\\d)")
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := reg.FindAllStringSubmatch(line, -1)
		first := digits[0][0]
		last := digits[len(digits)-1][0]
		number, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			panic(err)
		}
		total += number
	}
	return total
}

func day1_2() int {
	f, err := os.Open("inputs/day1.txt")
	numberMapping := map[string]int{
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
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	getDigit := func(v string) int {
		iv, err := strconv.Atoi(v)
		if err != nil {
			iv = numberMapping[v]
		}

		return iv
	}
	total := 0
	i := 1
	diff := 0
	for scanner.Scan() {
		line := scanner.Text()
		number := processIndex(line, getDigit)

		total += number
		i++
	}
	fmt.Println(diff)
	return total
}

func processIndex(line string, getDigit func(string) int) int {
	neededs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
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

	return number
}
