package maxConsecutiveOnes

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	tmp, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			res = max(tmp, res)
			tmp = 0
		}
	}
	res = max(tmp, res)
	return res
}
