package moveZeroes

func moveZeroes(nums []int) {
	cnt := 0
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
		} else {
			nums[k] = nums[i]
			k++
		}
	}
	for i := len(nums) - 1; i >= len(nums)-cnt; i-- {
		nums[i] = 0
	}
}
