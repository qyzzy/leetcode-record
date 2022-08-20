package candy

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	res := 0
	need := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	for i := len(ratings) - 1; i >= 1; i-- {
		if ratings[i] < ratings[i-1] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}
	for i := 0; i < len(need); i++ {
		res += need[i]
	}
	return res
}
