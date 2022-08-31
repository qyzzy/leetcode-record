package nonOverlappingIntervals

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Println(intervals)

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			res++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}
	return res
}
