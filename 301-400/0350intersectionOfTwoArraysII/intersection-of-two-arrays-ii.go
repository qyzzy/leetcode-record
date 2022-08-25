package intersectionOfTwoArraysII

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if e, ok := m[nums2[i]]; ok {
			if e != 0 {
				res = append(res, nums2[i])
				m[nums2[i]]--
			}
		}
	}
	return res
}
