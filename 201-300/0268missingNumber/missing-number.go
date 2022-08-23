package missingNumber

import "sort"

func missingNumber(nums []int) int {
	l := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != l {
			return l
		}
		l++
	}
	return len(nums)
}
