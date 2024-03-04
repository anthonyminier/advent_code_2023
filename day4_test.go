package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	result := day4("inputs/day4.txt")
	assert.Equal(t, 13, result)
}

func TestDay4_2(t *testing.T) {
	result := day4_2("inputs/day4_test.txt")
	assert.Equal(t, 30, result)
}
