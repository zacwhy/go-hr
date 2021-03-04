package q4

import (
	"testing"
)

var tests = []struct {
	before [][]int32
	after  [][]int32
}{
	{
		[][]int32{
			[]int32{1, 1},
			[]int32{2, 0},
		},
		[][]int32{
			[]int32{1, 2},
			[]int32{3, 4},
		},
	},
	{
		[][]int32{
			[]int32{2, 3},
			[]int32{5, 7},
		},
		[][]int32{
			[]int32{2, 5},
			[]int32{7, 17},
		},
	},
	{
		[][]int32{
			[]int32{1, 1, 1},
			[]int32{2, 0, 1},
			[]int32{1, 1, 1},
		},
		[][]int32{
			[]int32{1, 2, 3},
			[]int32{3, 4, 6},
			[]int32{4, 6, 9},
		},
	},
}

func TestFindAfterMatrix(t *testing.T) {
	for _, test := range tests {
		before := test.before
		want := test.after
		got := findAfterMatrix(before)

		if !matrixEqual(got, want) {
			t.Fatalf("findAfterMatrix(%v) = %v; want %v", before, got, want)
		}
	}
}

func TestFindBeforeMatrix(t *testing.T) {
	for _, test := range tests {
		after := test.after
		want := test.before
		got := findBeforeMatrix(after)

		if !matrixEqual(got, want) {
			t.Fatalf("findBeforeMatrix(%v) = %v; want %v", after, got, want)
		}
	}
}

func matrixEqual(a, b [][]int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range b {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range b[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func findBeforeMatrix(after [][]int32) [][]int32 {
	before := make([][]int32, len(after))

	for i := range after {
		before[i] = make([]int32, len(after[i]))

		for j := range after[i] {
			before[i][j] = findBefore(after, i, j)
		}
	}

	return before
}

func findBefore(after [][]int32, i, j int) int32 {
	s := after[i][j]

	if i > 0 {
		s -= after[i-1][j]
	}

	if j > 0 {
		s -= after[i][j-1]
	}

	if i > 0 && j > 0 {
		s += after[i-1][j-1]
	}

	return s
}

func findAfterMatrix(before [][]int32) [][]int32 {
	after := make([][]int32, len(before))

	for i := range before {
		after[i] = make([]int32, len(before[i]))

		for j := range before[i] {
			after[i][j] = findAfter(before, i, j)
		}
	}

	return after
}

func findAfter(before [][]int32, x, y int) int32 {
	var s int32

	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			s += before[i][j]
		}
	}

	return s
}
