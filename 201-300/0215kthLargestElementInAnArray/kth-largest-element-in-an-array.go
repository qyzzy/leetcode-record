package kthLargestElementInAnArray

func findKthLargest(nums []int, k int) int {
	b := make([]int, 20010)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			b[nums[i]+10000]++
		} else {
			b[-nums[i]]++
		}
	}
	i := 20009
	for i > 10000 {
		for b[i] > 0 {
			b[i]--
			k--
			if k == 0 {
				return i - 10000
			}
		}
		i--
	}
	i = 0
	for b[i] > 0 {
		b[i]--
		k--
		if k == 0 {
			return 0
		}
	}
	i = 1
	for i <= 10000 {
		for b[i] > 0 {
			b[i]--
			k--
			if k == 0 {
				return -i
			}
		}
		i++
	}
	return 0
}
