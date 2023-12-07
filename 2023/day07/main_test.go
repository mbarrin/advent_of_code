package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    hands
		expected hands
	}{
		"one": {
			input: hands{
				{cards: []int{13, 13, 6, 7, 7}},
				{cards: []int{13, 10, 11, 11, 10}},
			},
			expected: hands{
				{cards: []int{13, 10, 11, 11, 10}},
				{cards: []int{13, 13, 6, 7, 7}},
			},
		},
		"two": {
			input: hands{
				{cards: []int{12, 12, 12, 11, 14}},
				{cards: []int{10, 5, 5, 11, 5}},
			},
			expected: hands{
				{cards: []int{10, 5, 5, 11, 5}},
				{cards: []int{12, 12, 12, 11, 14}},
			},
		},
		"three": {
			input: hands{
				{cards: []int{14, 11, 14, 14, 14}},
				{cards: []int{8, 8, 12, 8, 8}},
				{cards: []int{14, 14, 14, 14, 8}},
				{cards: []int{10, 12, 12, 12, 12}},
				{cards: []int{12, 11, 12, 12, 12}},
				{cards: []int{13, 11, 13, 13, 13}},
				{cards: []int{14, 12, 12, 12, 12}},
				{cards: []int{10, 10, 12, 10, 10}},
				{cards: []int{14, 12, 14, 14, 14}},
				{cards: []int{9, 13, 9, 9, 9}},
			},
			expected: hands{
				{cards: []int{8, 8, 12, 8, 8}},
				{cards: []int{9, 13, 9, 9, 9}},
				{cards: []int{10, 10, 12, 10, 10}},
				{cards: []int{10, 12, 12, 12, 12}},
				{cards: []int{12, 11, 12, 12, 12}},
				{cards: []int{13, 11, 13, 13, 13}},
				{cards: []int{14, 11, 14, 14, 14}},
				{cards: []int{14, 12, 12, 12, 12}},
				{cards: []int{14, 12, 14, 14, 14}},
				{cards: []int{14, 14, 14, 14, 8}},
			},
		},
	}

	for name, tc := range tests {
		sort.Sort(tc.input)
		assert.Equal(t, tc.expected, tc.input, name)
	}
}
