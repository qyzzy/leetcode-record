package partitionEqualSubsetSum

import "sort"

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sort.Ints(nums)
	sum /= 2
	dp := make([]int, sum+1)
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	//fmt.Println(dp)
	return dp[sum] == sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
