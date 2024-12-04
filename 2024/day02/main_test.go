package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSafeTwo(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected bool
	}{
		//// input.txt
		//"one": {
		//	input:    []int{44, 51, 52, 54, 57, 60, 61, 63},
		//	expected: true,
		//},
		//"two": {
		//	input:    []int{11, 14, 15, 21, 18},
		//	expected: true,
		//},
		//// TODO: fix me
		//"three": {
		//	input:    []int{33, 35, 34, 35, 37, 40, 41},
		//	expected: true,
		//},
		//// sample.txt
		//"four": {
		//	input:    []int{7, 6, 4, 2, 1},
		//	expected: true,
		//},
		//"five": {
		//	input:    []int{1, 2, 7, 8, 9},
		//	expected: false,
		//},
		//"six": {
		//	input:    []int{9, 7, 6, 2, 1},
		//	expected: false,
		//},
		//"seven": {
		//	input:    []int{1, 3, 2, 4, 5},
		//	expected: true,
		//},
		//"eight": {
		//	input:    []int{8, 6, 4, 4, 1},
		//	expected: true,
		//},
		//"nine": {
		//	input:    []int{1, 3, 6, 7, 9},
		//	expected: true,
		//},
		//// More
		//"ten": {
		//	input:    []int{11, 9, 13, 15, 18, 21},
		//	expected: true,
		//},
		"eleven": {
			input:    []int{32, 30, 31, 32, 35},
			expected: true,
		},
	}

	for name, tc := range tests {
		output := isSafeTwo(tc.input)
		assert.Equal(t, tc.expected, output, name)
	}
}
