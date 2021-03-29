package hr

import (
	"testing"
)

func TestMinimumSwaps(t *testing.T) {
	tests := []struct {
		arr  []int32
		want int32
	}{
		{arr: []int32{7, 1, 3, 2, 4, 5, 6}, want: int32(5)},
		{arr: []int32{4, 3, 1, 2}, want: int32(3)},
		{arr: []int32{2, 3, 4, 1, 5}, want: int32(3)},
		{arr: []int32{1, 3, 5, 2, 4, 6, 7}, want: int32(3)},
	}
	for _, tt := range tests {
		got := minimumSwaps(tt.arr)
		if got != tt.want {
			t.Fatalf("minimumSwaps(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}

func minimumSwaps(arr []int32) int32 {
	swaps := int32(0)
	// fmt.Println(arr, "start")

	for i := 0; i < len(arr); {
		// check whether current value is in correct position
		if arr[i] == int32(i+1) {
			// move on to next index only when
			// the current index has the correct value
			i++
			continue
		}

		// do the swap
		j := arr[i] - 1
		arr[j], arr[i] = arr[i], arr[j]

		// fmt.Println(arr, "swap")

		swaps++
	}

	// fmt.Println(arr, "end")
	return swaps
}
