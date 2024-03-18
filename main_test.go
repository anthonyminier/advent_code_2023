package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNumberSuits(t *testing.T) {
	results := parseNumberSuits("av143 bhe129 zfe 032")
	assert.Equal(t, []int{143, 129, 32}, results)
	results = parseNumberSuits("	143	129	032")
	assert.Equal(t, []int{143, 129, 32}, results)

}
