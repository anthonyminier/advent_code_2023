package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	result := day2("inputs/day2.txt")
	assert.Equal(t, 2377, result)
}

func TestDay2_2(t *testing.T) {
	result := day2_2("inputs/day2.txt")
	assert.Equal(t, 71220, result)
}
