package medianOfTwoSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	tmp := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		val := 0
		if nums1[i] < nums2[j] {
			val = nums1[i]
			i++
		} else {
			val = nums2[j]
			j++
		}
		tmp = append(tmp, val)
	}
	for i < len(nums1) {
		tmp = append(tmp, nums1[i])
		i++
	}
	for j < len(nums2) {
		tmp = append(tmp, nums2[j])
		j++
	}
	//fmt.Println(tmp)
	if len(tmp)%2 == 0 {
		return float64(tmp[len(tmp)/2-1]+tmp[len(tmp)/2]) / 2
	} else {
		return float64(tmp[len(tmp)/2])
	}
	return 0.0
}
