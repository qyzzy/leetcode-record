package targetSum

func findTargetSumWays(nums []int, target int) int {
	res := 0
	var dfs func(p, sum int)
	dfs = func(p, sum int) {
		if p == len(nums) {
			if sum == target {
				res++
			}
			return
		}
		dfs(p+1, sum+nums[p])
		dfs(p+1, sum-nums[p])
	}
	dfs(0, 0)
	return res
}
