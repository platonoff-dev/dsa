package _0045_jump_game_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			nums: []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			want: 2,
		},
		{
			nums: []int{1, 1, 1, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			got := jump(tt.nums)
			assert.Equal(t, tt.want, got, "input: %v", tt.nums)
		})
	}
}
