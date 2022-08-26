package kthSmallestElementInASortedMatrix

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	nums := make([]int, m*n)
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nums[index] = matrix[i][j]
			index++
		}
	}
	sort.Ints(nums)
	return nums[k-1]
}
