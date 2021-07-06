package lc

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	nums   []int
	target int
	want   []int
}{
	{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
	{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
	{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		name := fmt.Sprintf("TwoSum(%v, %v)", tt.nums, tt.target)
		t.Run(name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}

	for j, num := range nums {
		diff := target - num

		if i, ok := m[diff]; ok {
			return []int{i, j}
		}

		m[num] = j
	}

	return nil
}

func TwoSumBrute(nums []int, target int) []int {
	count := len(nums)
	for i := range nums {
		for j := i + 1; j < count; j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
