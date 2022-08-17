// Source: https://leetcode.cn/problems/two-sum/
// Author: qyzzy

package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if tmp, ok := m[target-v]; ok {
			return []int{tmp, i}
		}
		m[v] = i
	}
	return nil
}
