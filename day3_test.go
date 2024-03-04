package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	result := day3("inputs/day3.txt")
	assert.Equal(t, 550934, result)
}

func TestDay3_2(t *testing.T) {
	result := day3_2("inputs/day3.txt")
	assert.Equal(t, 81997870, result)
}
