package wiggleSortII

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	t := make([]int, n)
	copy(t, nums)
	sort.Ints(t)
	j, k := n, (n+1)/2
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			k--
			nums[i] = t[k]
		} else {
			j--
			nums[i] = t[j]
		}
	}
}
