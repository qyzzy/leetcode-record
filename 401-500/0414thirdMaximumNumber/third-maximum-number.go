package thirdMaximumNumber

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	rank := 1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] == nums[i-1] {
			continue
		} else {
			rank++
		}
		if rank == 3 {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}
