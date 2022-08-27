package wiggleSubsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	cur, pre := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i+1] - nums[i]
		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0) {
			res++
			pre = cur
		}
	}
	return res
}
