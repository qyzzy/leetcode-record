package maxPointsOnALine

func maxPoints(points [][]int) int {
	n := len(points)
	res := 1
	for i := 0; i < n; i++ {
		x := points[i]
		for j := i + 1; j < n; j++ {
			y := points[j]
			cnt := 2
			for k := j + 1; k < n; k++ {
				z := points[k]
				s1 := (x[0] - y[0]) * (y[1] - z[1])
				s2 := (x[1] - y[1]) * (y[0] - z[0])
				if s1 == s2 {
					cnt++
				}
			}
			res = max(res, cnt)
		}
	}
	return res
}
func max(args ...int) int {
	res := args[0]
	for _, item := range args {
		if item > res {
			res = item
		}
	}
	return res
}
