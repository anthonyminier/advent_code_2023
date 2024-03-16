package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := day5_2()

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
