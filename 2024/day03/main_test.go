package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveMultsRegex(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"one": {
			input:    "8",
			expected: "8",
		},
		"two": {
			input:    "8don't()9do()",
			expected: "8",
		},
		"three": {
			input:    "8don't()9do()10do()11",
			expected: "810do()11",
		},
		"four": {
			input:    "7don't()8don't()9do()10do()",
			expected: "710do()",
		},
		"five": {
			input:    "7don't()8do()9don't()10",
			expected: "79",
		},
		"six": {
			input:    "don't()9don't()10do()3",
			expected: "3",
		},
	}

	for name, tc := range tests {
		output := removeMults([]byte(tc.input))
		assert.Equal(t, []byte(tc.expected), output, name)
	}
}
