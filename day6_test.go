package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	result := day6("inputs/day6.txt")
	assert.Equal(t, 6209190, result)
}

// this test timeout due to around 2 minutes to run this bruteforce method, there is probably
// a smarter solution but i don't know which one
func TestDay6_2(t *testing.T) {
	result := day6_2("inputs/day6_2.txt")
	assert.Equal(t, 28545089, result)
}
