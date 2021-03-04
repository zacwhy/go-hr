package q3

import (
	"testing"
)

func TestMaxDifference(t *testing.T) {
	var tests = []struct {
		in   []int32
		want int32
	}{
		{[]int32{2, 3, 10, 2, 4, 8, 1}, 8},
		{[]int32{7, 9, 5, 6, 3, 2}, 2},
		{[]int32{3}, -1},
		{[]int32{-3, -2}, 1},
		{[]int32{}, -1},
	}

	for _, test := range tests {
		in := test.in
		want := test.want
		got := maxDifference(in)

		if got != want {
			t.Errorf("maxDifference(%v) = %v; want %v", in, got, want)
		}
	}
}

func maxDifference(px []int32) int32 {
	maxDiff := int32(-1)
	var min int32

	for i, v := range px {
		if i == 0 {
			min = v
			continue
		}

		if diff := v - min; diff > maxDiff {
			maxDiff = diff
		}

		if v < min {
			min = v
		}
	}

	return maxDiff
}
