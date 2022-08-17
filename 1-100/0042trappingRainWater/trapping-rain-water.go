package trappingRainWater

func trap(height []int) int {
	n := len(height)
	ld, rd := make([]int, n), make([]int, n)
	lmax, rmax := 0, 0
	for i := 0; i < n; i++ {
		lmax = max(lmax, height[i])
		ld[i] = max(ld[i], lmax)
	}
	for i := n - 1; i >= 0; i-- {
		rmax = max(rmax, height[i])
		rd[i] = max(rd[i], rmax)
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(ld[i], rd[i]) - height[i]
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
