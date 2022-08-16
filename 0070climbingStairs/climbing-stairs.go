package climbingStairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	p1, p2 := 2, 1
	res := 0
	for i := 3; i <= n; i++ {
		res = p1 + p2
		p2 = p1
		p1 = res
	}
	return res
}
