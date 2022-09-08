package minimumMovesToEqualArrayElementsII

import "sort"

func minMoves2(nums []int) int {
	res := 0
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		res += abs(num - x)
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
