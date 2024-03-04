package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	result := day1()
	assert.Equal(t, 53974, result)
}

func TestDay1_2(t *testing.T) {
	result := day1_2()
	assert.Equal(t, 52840, result)
}
