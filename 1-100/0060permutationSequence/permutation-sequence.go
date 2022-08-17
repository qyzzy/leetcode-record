package permutationSequence

func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = byte(i + '0')
	}
	return string(helper(nums, k))
}

func helper(nums []byte, k int) []byte {
	if len(nums) == 2 {
		if k == 1 {
			return nums
		} else {
			nums[0], nums[1] = nums[1], nums[0]
			return nums
		}
	}
	base := 1
	for i := 1; i < len(nums); i++ {
		base *= i
	}
	var res []byte
	for i := 0; i < len(nums); i++ {
		if k > base {
			k -= base
			continue
		}
		res = append(res, nums[i])
		nums = append(nums[:i], nums[i+1:]...)
		res = append(res, helper(nums, k)...)
		return res
	}
	return res
}
