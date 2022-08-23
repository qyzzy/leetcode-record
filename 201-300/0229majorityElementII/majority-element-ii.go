package majorityElementII

func majorityElement(nums []int) []int {
	n := len(nums)
	res := []int{}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v > n/3 {
			res = append(res, k)
		}
	}
	return res
}
