package main

import "slices"

func hIndex(citations []int) int {
	slices.Sort(citations)
	for i := 0; i < len(citations); i++ {
		if citations[len(citations)-1-i] < i+1 {
			return i
		}
	}
	return len(citations)
}
