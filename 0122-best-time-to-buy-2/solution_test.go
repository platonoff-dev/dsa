package _122_best_time_to_buy_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.prices), func(t *testing.T) {
			got := maxProfit(tt.prices)
			assert.Equal(t, tt.want, got, "input: %v", tt.prices)
		})
	}
}
