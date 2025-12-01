package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := map[string]struct {
		start            int
		count            int
		expectedLocation int
		expectedLoops    int
	}{
		"loop once right": {
			start:            45,
			count:            60,
			expectedLocation: 5,
			expectedLoops:    1,
		},
		"loops 0 time right from 0": {
			start:            0,
			count:            1,
			expectedLocation: 1,
			expectedLoops:    0,
		},
		"loops 1 time right starting and ending at 0": {
			start:            0,
			count:            200,
			expectedLocation: 0,
			expectedLoops:    1,
		},
		"loop 5 times right": {
			start:            33,
			count:            503,
			expectedLocation: 36,
			expectedLoops:    5,
		},
		"sample 2": {
			start:            95,
			count:            60,
			expectedLocation: 55,
			expectedLoops:    1,
		},
		"loop once left": {
			start:            51,
			count:            -60,
			expectedLocation: 91,
			expectedLoops:    1,
		},
		"loops 2 times left": {
			start:            6,
			count:            -205,
			expectedLocation: 1,
			expectedLoops:    2,
		},
		"loops 3 times left": {
			start:            1,
			count:            -202,
			expectedLocation: 99,
			expectedLoops:    3,
		},
		"loops 0 time left from 0": {
			start:            0,
			count:            -1,
			expectedLocation: 99,
			expectedLoops:    0,
		},
		"loops 1 time left starting and ending at 0": {
			start:            0,
			count:            -200,
			expectedLocation: 0,
			expectedLoops:    1,
		},
		"sample 1": {
			start:            50,
			count:            -68,
			expectedLocation: 82,
			expectedLoops:    1,
		},
	}
	for name, tc := range tests {
		location, loops := rotate(tc.start, tc.count)
		assert.Equal(t, tc.expectedLocation, location, name)
		assert.Equal(t, tc.expectedLoops, loops, name)
	}
}
