package insertInterval

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = max(intervals[i][1], intervals[i-1][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}
	return intervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
