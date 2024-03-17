package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	result := day5()
	assert.Equal(t, 111627841, result)
}

// this test timeout due to around 2 minutes to run this bruteforce method, there is probably
// a smarter solution but i don't know which one
func TestDay5_2(t *testing.T) {
	result := day5_2()
	assert.Equal(t, 69323688, result)
}
