package triangle

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
	}
	res := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
			if i == len(triangle)-1 {
				res = min(res, triangle[i][j])
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
