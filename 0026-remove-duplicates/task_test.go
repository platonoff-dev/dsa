package _026_remove_duplicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input       []int
		expectedN   int
		expectedArr []int
	}{
		{
			input:       []int{1, 1, 2},
			expectedN:   2,
			expectedArr: []int{1, 2},
		},
		{
			input:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedN:   5,
			expectedArr: []int{0, 1, 2, 3, 4},
		},
		{
			input:       []int{1, 1},
			expectedN:   1,
			expectedArr: []int{1},
		},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprint(c.input), func(t *testing.T) {
			result := removeDuplicates(c.input)
			assert.Equal(t, c.expectedN, result)
			assert.Equal(t, c.expectedArr, c.input[:c.expectedN])
		})
	}
}
