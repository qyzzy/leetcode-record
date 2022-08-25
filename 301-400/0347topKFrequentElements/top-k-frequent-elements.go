package topKFrequentElements

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	un := make([]int, len(m))
	ind := 0
	for v := range m {
		un[ind] = v
		ind++
	}
	sort.Slice(un, func(i, j int) bool {
		a, b := un[i], un[j]
		return m[a] > m[b] || m[a] == m[b] && a < b
	})
	return un[:k]
}
