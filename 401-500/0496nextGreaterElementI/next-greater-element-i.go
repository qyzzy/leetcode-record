package nextGreaterElementI

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	tmp := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		tmp[nums2[i]] = -1
		for j := i; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				tmp[nums2[i]] = nums2[j]
				break
			}
		}
	}
	for i := 0; i < len(nums1); i++ {
		res[i] = tmp[nums1[i]]
	}
	return res
}
