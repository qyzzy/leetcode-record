package sortColors

func sortColors(nums []int) {
	mp := [3]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	i := 0
	for j := 0; j < 3; j++ {
		for mp[j] > 0 {
			nums[i] = j
			mp[j]--
			i++
		}
	}
}
