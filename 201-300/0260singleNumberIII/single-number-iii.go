package singleNumberIII

func singleNumber(nums []int) []int {
	m := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
