package arithmeticSlices

func numberOfArithmeticSlices(nums []int) int {
	res, l := 0, 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]-nums[i-1] == nums[i+1]-nums[i] {
			l++
		} else {
			l = 0
		}
		res += l
	}
	return res
}
