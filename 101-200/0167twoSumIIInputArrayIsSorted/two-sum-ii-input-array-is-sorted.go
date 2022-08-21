package twoSumIIInputArrayIsSorted

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

//func twoSum(numbers []int, target int) []int {
//	m := map[int]int{}
//	res := make([]int, 2)
//	for i := 0; i < len(numbers); i++ {
//		m[numbers[i]] = i
//	}
//	for i := 0; i < len(numbers); i++ {
//		if _, ok := m[target - numbers[i]]; ok {
//			if m[target - numbers[i]] != i {
//				res[0] = i + 1
//				res[1] = m[target - numbers[i]] + 1
//				break
//			}
//		}
//	}
//	return res
//}
