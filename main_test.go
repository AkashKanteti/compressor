package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TextParsing(t *testing.T) {
	str := "all is well"
	expectedMap := map[rune]int{'a': 1, 'l': 4, ' ': 2, 'i': 1, 's': 1, 'w': 1, 'e': 1}
	actualMap := parsing(str)
	assert.Equal(t, expectedMap, actualMap)
}
