package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := day5_2("inputs/day5.txt")

	fmt.Println(result)
}

func fromInput(input string, each func(line string)) {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		each(scanner.Text())
	}
}

func parseNumberSuits(str string) []int {
	numbers := []int{}
	currentNumber := ""
	for _, c := range str {
		if strings.Contains("0123456789", string(c)) {
			currentNumber += string(c)
			continue
		}
		if currentNumber != "" {
			num, _ := strconv.Atoi(currentNumber)
			numbers = append(numbers, num)
			currentNumber = ""
		}
	}
	if currentNumber != "" {
		num, _ := strconv.Atoi(currentNumber)
		numbers = append(numbers, num)
	}

	return numbers
}
