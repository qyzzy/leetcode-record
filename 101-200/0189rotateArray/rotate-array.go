package rotateArray

func rotate(nums []int, k int) {
	k = k % len(nums)
	k = len(nums) - k
	newNums := append(nums[k:], nums[:k]...)
	copy(nums, newNums)
}
