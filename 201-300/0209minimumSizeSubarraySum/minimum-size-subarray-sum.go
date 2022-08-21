package minimumSizeSubarraySum

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	l, r, sum := 0, 0, 0
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			if res > r-l+1 {
				res = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
