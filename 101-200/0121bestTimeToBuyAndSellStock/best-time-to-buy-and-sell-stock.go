package bestTimeToBuyAndSellStock

func maxProfit(nums []int) int {
	res := 0
	minP := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i]-minP)
		minP = min(nums[i], minP)
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
