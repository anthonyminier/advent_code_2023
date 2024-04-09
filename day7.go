package main

import (
	"sort"
	"strconv"
	"strings"
)

const (
	high    = 1
	onePair = 2
	twoPair = 3
	three   = 4
	full    = 5
	four    = 6
	five    = 7
)

var cardStrength = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardStrengthJoker = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type hand struct {
	cards string
	bid   int
	hand  int
}

func getHand(hm map[string]int) int {
	// five of a kind (AAAAA)
	if len(hm) == 1 {
		return five
	}
	// four of a kind (AA8AA) or full (AA8A8)
	if len(hm) == 2 {
		for _, nbOccurences := range hm {
			if nbOccurences == 4 || nbOccurences == 1 {
				return four
			}
			return full
		}
	}
	// three of a kind (AA87A)
	if len(hm) == 3 {
		for _, nbOccurences := range hm {
			if nbOccurences == 3 {
				return three
			}
		}
		return twoPair
	}
	if len(hm) == 4 {
		return onePair
	}
	return high
}

func day7(input string) int {
	hs := []hand{}
	// format the data
	fromInput(input, func(line string) {
		args := strings.Split(line, " ")
		bid, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		cards := args[0]
		handMapping := map[string]int{}
		for _, card := range cards {
			nbOccurences, _ := handMapping[string(card)]
			handMapping[string(card)] = nbOccurences + 1
		}

		hs = append(hs, hand{cards: args[0], bid: bid, hand: getHand(handMapping)})
	})

	sort.Slice(hs, func(i, j int) bool {
		if hs[i].hand == hs[j].hand {
			for index := range 5 {
				if hs[i].cards[index] == hs[j].cards[index] {
					continue
				}
				return cardStrength[hs[i].cards[index]] < cardStrength[hs[j].cards[index]]
			}
		}

		return hs[i].hand < hs[j].hand
	})

	total := 0
	for i, h := range hs {
		total += ((i + 1) * h.bid)
	}

	return total
}

func getHandWithJoker(hm map[string]int) int {
	if len(hm) == 1 {
		return five
	}

	// get nb of jokers
	nbJoker := hm["J"]
	if nbJoker == 0 {
		return getHand(hm)
	}

	switch nbJoker {
	case 4:
		return five
	case 3:
		if len(hm) == 2 {
			return five
		}
		return four
	case 2:
		if len(hm) == 2 {
			return five
		}
		if len(hm) == 3 {
			return four
		}
		if len(hm) == 4 {
			return three
		}
	case 1:
		if len(hm) == 2 {
			return five
		}
		if len(hm) == 3 {
			maxOccurence := 1
			for _, nbOccurence := range hm {
				if nbOccurence > maxOccurence {
					maxOccurence = nbOccurence
				}
			}
			if maxOccurence == 3 {
				return four
			}
			return full
		}
		if len(hm) == 4 {
			return three
		}
		return onePair
	}

	return high
}

func day7_2(input string) int {
	hs := []hand{}
	// format the data
	fromInput(input, func(line string) {
		args := strings.Split(line, " ")
		bid, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		cards := args[0]
		handMapping := map[string]int{}
		for _, card := range cards {
			nbOccurences, _ := handMapping[string(card)]
			handMapping[string(card)] = nbOccurences + 1
		}

		hs = append(hs, hand{cards: args[0], bid: bid, hand: getHandWithJoker(handMapping)})
	})

	sort.Slice(hs, func(i, j int) bool {
		if hs[i].hand == hs[j].hand {
			for index := range 5 {
				if hs[i].cards[index] == hs[j].cards[index] {
					continue
				}
				return cardStrengthJoker[hs[i].cards[index]] < cardStrengthJoker[hs[j].cards[index]]
			}
		}

		return hs[i].hand < hs[j].hand
	})

	total := 0
	for i, h := range hs {
		total += ((i + 1) * h.bid)
	}

	return total
}
