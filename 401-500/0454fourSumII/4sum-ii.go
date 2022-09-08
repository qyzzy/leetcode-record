package _4sumII

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	m := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			//v1 + v2 + v3 + v4
			res += m[-v3-v4]
		}
	}
	return res
}
