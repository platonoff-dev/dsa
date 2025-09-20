package _0274_h_index

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		want      int
	}{
		{
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
		{
			citations: []int{1, 3, 1},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.citations), func(t *testing.T) {
			got := hIndex(tt.citations)
			assert.Equal(t, tt.want, got, "input: %v", tt.citations)
		})
	}
}
