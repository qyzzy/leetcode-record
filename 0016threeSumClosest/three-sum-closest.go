package threeSumClosest

import "sort"

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func threeSumClosest(nums []int, target int) int {
	res := -100000
	l, r := 0, 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
		l, r = i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			//fmt.Println(sum)
			if sum == target {
				return target
			}
			if sum > target {
				if res == -100000 {
					res = sum
				}
				if abs(res-target) > abs(sum-target) {
					res = sum
				}
				for r-1 > l && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
			if sum < target {
				if res == -100000 {
					res = sum
				}
				if abs(res-target) > abs(sum-target) {
					res = sum
				}
				for l+1 < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
	}
	return res
}
