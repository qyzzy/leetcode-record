package findAllDuplicatesInAnArray

func findDuplicates(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		x := (nums[i] - 1) % len(nums)
		nums[x] += len(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > len(nums)*2 {
			res = append(res, i+1)
		}
	}
	return res
}
