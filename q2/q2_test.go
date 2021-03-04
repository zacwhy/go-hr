package q2

import (
	"sort"
	"testing"
)

func TestGetMinimumMoves(t *testing.T) {
	var tests = []struct {
		in   []int32
		want int32
	}{
		{[]int32{5, 1, 3, 2}, 2},             // 3 -> 5
		{[]int32{1, 3, 2, 5}, 2},             // 3 -> 5
		{[]int32{5, 2, 1, 3}, 3},             // 2 -> 3 -> 5
		{[]int32{1, 3, 5, 2}, 2},             // 3 -> 5
		{[]int32{3, 1, 3, 5, 2}, 3},          // 3 -> 3 -> 5
		{[]int32{8, 6, 1, 7, 5, 2, 3, 9}, 5}, // 5 -> 6 -> 7 -> 8 -> 9
	}

	for _, test := range tests {
		in := test.in
		want := test.want
		got := getMinimumMoves(in)

		if got != want {
			t.Errorf("getMinimumMoves(%v) = %v; want %v", in, got, want)
		}
	}
}

func getMinimumMoves(arr []int32) int32 {
	sortedArr := make([]int, len(arr))
	for i, n := range arr {
		sortedArr[i] = int(n)
	}
	sort.Ints(sortedArr)

	j := 0

	for _, n := range arr {
		if n == int32(sortedArr[j]) {
			j++
		}
	}

	return int32(len(arr) - j)
}
