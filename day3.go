package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regDigit = regexp.MustCompile("^\\d$")

func day3(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}
	total := 0
	for j := range len(board) {
		for i := 0; i < len(board[j]); i++ {
			if !regDigit.MatchString(board[j][i]) {
				continue
			}
			n, hasSiblingSymbal := identifySiblings(board, j, i, len(board), len(board[j]))
			if hasSiblingSymbal {
				d, _ := strconv.Atoi(n)
				total += d
			}
			i += len(n) - 1
		}
	}

	return total
}

func identifySiblings(board [][]string, j, i, jlength, ilength int) (string, bool) {
	n := board[j][i]
	hasSiblingSymbol := isNotADot(board, j-1, i-1, jlength, ilength) ||
		isNotADot(board, j-1, i, jlength, ilength) ||
		isNotADot(board, j-1, i+1, jlength, ilength) ||
		isNotADot(board, j, i-1, jlength, ilength) ||
		isNotADot(board, j, i+1, jlength, ilength) ||
		isNotADot(board, j+1, i-1, jlength, ilength) ||
		isNotADot(board, j+1, i, jlength, ilength) ||
		isNotADot(board, j+1, i+1, jlength, ilength)

	// recursive if the digit has another digit on its right
	if i+1 < ilength && regDigit.MatchString(board[j][i+1]) {
		neighbourDigit, hasNeighbourSiblingSymbol := identifySiblings(board, j, i+1, jlength, ilength)
		return fmt.Sprintf("%s%s", n, neighbourDigit), cmp.Or(hasSiblingSymbol, hasNeighbourSiblingSymbol)
	}

	return n, hasSiblingSymbol
}

func isNotADot(board [][]string, j, i, jlength, ilength int) bool {
	return j >= 0 && i >= 0 && j < jlength && i < ilength && !regDigit.MatchString(board[j][i]) && board[j][i] != "."
}

func day3_2(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	board := [][]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}
	total := 0
	numberCoordinates := map[string]string{}
	starCoordinates := []string{}
	for j := range len(board) {
		for i := 0; i < len(board[j]); i++ {
			if board[j][i] == "*" {
				starCoordinates = append(starCoordinates, fmt.Sprintf("%d-%d", j, i))
				continue
			}
			if !regDigit.MatchString(board[j][i]) {
				continue
			}
			n, _ := identifySiblings(board, j, i, len(board), len(board[j]))
			numberCoordinates[fmt.Sprintf("%d-%d", j, i)] = n
			i += len(n) - 1
		}
	}
	totalStarSiblings := 0
	for _, sCoord := range starCoordinates {
		sCoords := strings.Split(sCoord, "-")
		o, _ := strconv.Atoi(sCoords[0])
		x, _ := strconv.Atoi(sCoords[1])
		minO := o - 1
		minX := x - 1
		maxO := o + 1
		maxX := x + 1
		numbersSibling := []string{}
		for coord, number := range numberCoordinates {
			ords := strings.Split(coord, "-")
			on, _ := strconv.Atoi(ords[0])
			xn, _ := strconv.Atoi(ords[1])
			isSiblings := false
			maxNumberX := xn + len(number)
			for ; xn < maxNumberX; xn++ {
				if on >= minO && on <= maxO && xn >= minX && xn <= maxX {
					isSiblings = true
					break
				}
			}
			if isSiblings {
				numbersSibling = append(numbersSibling, number)
			}
		}
		if len(numbersSibling) == 2 {
			n1, _ := strconv.Atoi(numbersSibling[0])
			n2, _ := strconv.Atoi(numbersSibling[1])
			total += (n1 * n2)
			totalStarSiblings++
		}
	}

	return total
}
