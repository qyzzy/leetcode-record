package containerWithMostWater

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		res = max(res, (r-l)*min(height[r], height[l]))
		if height[r] < height[l] {
			r--
		} else {
			l++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
