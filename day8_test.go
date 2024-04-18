package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	result := day8("inputs/day8.txt")
	assert.Equal(t, 19783, result)
}

func TestDay8_2(t *testing.T) {
	result := day8_2("inputs/day8.txt")
	assert.Equal(t, 6, result)
}
