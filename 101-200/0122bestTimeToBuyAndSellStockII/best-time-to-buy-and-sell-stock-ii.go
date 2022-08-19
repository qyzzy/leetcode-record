package bestTimeToBuyAndSellStockII

func maxProfit(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res += nums[i] - nums[i-1]
		}
	}
	return res
}
