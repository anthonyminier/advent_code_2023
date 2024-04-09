package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	result := day7("inputs/day7.txt")
	// the code works correctly but the input changed between december and now...
	assert.Equal(t, 253603890, result)
}

// this test timeout due to around 2 minutes to run this bruteforce method, there is probably
// a smarter solution but i don't know which one
func TestDay7_2(t *testing.T) {
	result := day7_2("inputs/day7.txt")
	assert.Equal(t, 253630098, result)
}
