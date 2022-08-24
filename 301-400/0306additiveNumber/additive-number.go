package additiveNumber

import "strconv"

func isAdditiveNumber(num string) bool {
	if len(num) <= 2 {
		return false
	}
	var res bool
	var dfs func(p int, t []int)
	dfs = func(p int, t []int) {
		if p >= len(num) {
			if len(t) >= 3 {
				res = true
			}
			return
		}
		for i := p; i < len(num); i++ {
			if num[p] == '0' && p != i {
				return
			}
			n, _ := strconv.Atoi(num[p : i+1])
			if len(t) < 2 {
				t = append(t, n)
				dfs(i+1, t)
				t = t[:len(t)-1]
			} else {
				if n == t[len(t)-1]+t[len(t)-2] {
					t = append(t, n)
					dfs(i+1, t)
					t = t[:len(t)-1]
				}
			}
		}
	}
	dfs(0, []int{})
	return res
}
