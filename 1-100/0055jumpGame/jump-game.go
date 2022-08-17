package jumpGame

func canJump(nums []int) bool {
	most := 0
	for i := 0; i < len(nums); i++ {
		if most < i {
			return false
		}
		most = max(most, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
