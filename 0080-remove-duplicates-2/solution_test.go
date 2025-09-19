package _0080_remove_duplicates_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		wantK    int
		wantNums []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			wantK:    5,
			wantNums: []int{1, 1, 2, 2, 3},
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantK:    7,
			wantNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			result := removeDuplicates(tt.nums)
			assert.Equal(t, tt.wantK, result)
			assert.Equal(t, tt.wantNums, tt.nums[:result])
		})
	}
}
