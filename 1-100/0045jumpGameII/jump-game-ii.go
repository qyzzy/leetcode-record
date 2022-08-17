package jumpGameII

func jump(nums []int) int {
	res := 0
	border := 0
	maxD := 0
	for i := 0; i < len(nums)-1; i++ {
		maxD = max(maxD, nums[i]+i)
		if i == border {
			border = maxD
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
