package findTheDuplicateNumber

func findDuplicate(nums []int) int {
	b := map[int]int{}
	for i := 0; i < len(nums); i++ {
		b[nums[i]]++
		if b[nums[i]] >= 2 {
			return nums[i]
		}
	}
	return 0
}
