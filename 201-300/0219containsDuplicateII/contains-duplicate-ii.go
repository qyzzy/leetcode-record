package containsDuplicateII

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if prev, ok := m[nums[i]]; ok {
			if i-prev <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}
