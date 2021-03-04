package main

import (
	"testing"
)

func TestMinCost(t *testing.T) {
	var tests = []struct {
		numProjects int32
		projectId   []int32
		bid         []int32
		want        int64
	}{
		{2, []int32{0, 1, 0, 1, 1}, []int32{4, 74, 47, 744, 7}, 11},
	}

	for _, test := range tests {
		numProjects := test.numProjects
		projectId := test.projectId
		bid := test.bid
		want := test.want
		got := minCost(numProjects, projectId, bid)

		if got != want {
			t.Errorf("minCost(%v, %v, %v) = %v; want %v", numProjects, projectId, bid, got, want)
		}
	}
}

func minCost(numProjects int32, projectId []int32, bid []int32) int64 {
	lowestBids := make(map[int32]int32)

	for i, pid := range projectId {
		lowestBid, ok := lowestBids[pid]

		if !ok || bid[i] < lowestBid {
			lowestBids[pid] = bid[i]
		}
	}

	var cost int64

	for i := int32(0); i < numProjects; i++ {
		lowestBid, ok := lowestBids[i]

		if !ok {
			return -1
		}

		cost += int64(lowestBid)
	}

	return cost
}
