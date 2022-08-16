package removeDuplicatesFromSortedArrayII

func removeDuplicates(nums []int) int {
	c := 1
	for i := 1; i < len(nums); {
		if nums[i-1] == nums[i] {
			c++
			if c > 2 {
				nums = append(nums[:i-1], nums[i:]...)
				c = 2
			} else {
				i++
			}
		} else {
			i++
			c = 1
		}
	}
	return len(nums)
}
