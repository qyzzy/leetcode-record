package nextPermutation

import "sort"

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for i-1 >= 0 && nums[i-1] >= nums[i] {
		i--
	}
	i--
	if i == -1 {
		sort.Ints(nums)
		return
	}
	j := len(nums) - 1
	for j > i && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
