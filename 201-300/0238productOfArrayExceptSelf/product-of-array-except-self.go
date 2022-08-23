package productOfArrayExceptSelf

func productExceptSelf(nums []int) []int {
	left, right := 1, 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = left
		left *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
