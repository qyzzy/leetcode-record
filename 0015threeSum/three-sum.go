package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	mid := 0
	for l < r-1 {
		if l > 0 && nums[l-1] == nums[l] {
			l++
			continue
		}
		r = len(nums) - 1
		mid = l + 1
		for mid < r {
			sum := nums[l] + nums[r] + nums[mid]
			if sum == 0 {
				res = append(res, []int{nums[l], nums[mid], nums[r]})
				for mid+1 < r && nums[mid+1] == nums[mid] {
					mid++
				}
				for mid < r-1 && nums[r-1] == nums[r] {
					r--
				}
				mid++
				r--
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				mid++
			}
		}
		l++
	}
	return res
}
