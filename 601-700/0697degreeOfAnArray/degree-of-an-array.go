package degreeOfAnArray

import "math"

func findShortestSubArray(nums []int) int {
	maxN := []int{}
	d := 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > d {
			d = m[nums[i]]
			maxN = []int{nums[i]}
		} else if m[nums[i]] == d {
			maxN = append(maxN, nums[i])
		}
	}
	f := make([]int, len(maxN))
	for i := 0; i < len(maxN); i++ {
		f[i] = -1
	}
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(maxN); j++ {
			if nums[i] == maxN[j] {
				if f[j] == -1 {
					f[j] = i
				}
				m[nums[i]]--
				if m[nums[i]] == 0 {
					res = min(res, i-f[j]+1)
				}
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
