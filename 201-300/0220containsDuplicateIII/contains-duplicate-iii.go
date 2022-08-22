package containsDuplicateIII

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums) && i <= k; i++ {
		for j := i + 1; j < len(nums) && j <= k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	if len(nums)-1 < k {
		return false
	}
	for i := 1; i+k < len(nums); i++ {
		if check(nums[i:i+k+1], t) {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func check(a []int, t int) bool {
	key := a[len(a)-1]
	for i := 0; i < len(a)-1; i++ {
		if abs(key-a[i]) <= t {
			return true
		}
	}
	return false
}
