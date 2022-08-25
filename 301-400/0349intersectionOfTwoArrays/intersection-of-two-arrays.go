package intersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = i
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(m, nums2[i])
		}
	}
	return res
}
