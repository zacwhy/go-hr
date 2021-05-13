package lc

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{heights: []int{2}, want: 0},
		{heights: []int{2, 2}, want: 0},
		{heights: []int{1, 2, 1}, want: 0},
		{heights: []int{1, 1, 2, 1}, want: 0},
		{heights: []int{2, 0, 2}, want: 2},
		{heights: []int{2, 0, 1}, want: 1},
		{heights: []int{2, 0, 1, 1}, want: 1},
		{heights: []int{2, 0, 1, 2}, want: 3},
		{heights: []int{2, 1, 0, 1, 2}, want: 4},
	}
	for _, tt := range tests {
		got := trap(tt.heights)
		if got != tt.want {
			t.Errorf("fn(%v) = %v; want %v", tt.heights, got, tt.want)
		}
	}
}

func trap(heights []int) int {
	water := 0
	left := 0
	right := len(heights) - 1
	leftMax, rightMax := 0, 0

	for left < right {
		// move pointer of the lower side
		// so we can be sure water will be trapped when height at current index is lower than max height
		if heights[left] < heights[right] {
			if heights[left] >= leftMax {
				leftMax = heights[left]
			} else {
				water += leftMax - heights[left]
			}
			left++
		} else {
			if heights[right] >= rightMax {
				rightMax = heights[right]
			} else {
				water += rightMax - heights[right]
			}
			right--
		}
	}

	return water
}
