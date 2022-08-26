package lexicographicalNumbers

func lexicalOrder(n int) []int {
	res := []int{}
	var dfs func(t int)
	dfs = func(t int) {
		//fmt.Println(t, res)
		if t > n {
			return
		}
		res = append(res, t)
		for i := 0; i <= 9; i++ {
			if t == -1 {
				if i == 0 {
					continue
				} else {
					t = i
					dfs(t)
					t = -1
				}
			} else {
				t = t*10 + i
				if t > n {
					return
				}
				dfs(t)
				t = (t - i) / 10
			}
		}
	}
	dfs(-1)
	res = res[1:]
	return res
}
