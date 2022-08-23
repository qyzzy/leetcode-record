package hIndex

import "sort"

func hIndex(citations []int) int {
	h := 0
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}
