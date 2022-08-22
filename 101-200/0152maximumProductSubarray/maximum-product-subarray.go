package maximumProductSubarray

func maxProduct(nums []int) int {
	maxP, minP, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxP, minP
		maxP = max(nums[i], max(mn*nums[i], mx*nums[i]))
		minP = min(nums[i], min(mn*nums[i], mx*nums[i]))
		res = max(res, maxP)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
