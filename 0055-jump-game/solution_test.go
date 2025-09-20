package _0055_jump_game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			nums: []int{0, 1},
			want: false,
		},
		{
			nums: []int{0},
			want: true,
		},
		{
			nums: []int{2, 5, 0, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			got := canJump(tt.nums)
			assert.Equal(t, tt.want, got, "input: %v", tt.nums)
		})
	}
}
