package permutationsII

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	sort.Ints(nums)
	var dfs func(t []int, m map[int]int)
	dfs = func(t []int, m map[int]int) {
		if len(t) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			if m[nums[i]] <= 0 {
				continue
			}
			m[nums[i]]--
			t = append(t, nums[i])
			dfs(t, m)
			m[nums[i]]++
			t = t[:len(t)-1]
		}
	}
	dfs([]int{}, m)
	return res
}
