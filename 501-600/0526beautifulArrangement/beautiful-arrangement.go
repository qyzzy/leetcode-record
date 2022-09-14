package beautifulArrangement

func countArrangement(n int) int {
	count := 0
	m := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		m[i]++
	}
	var dfs func(i int, m map[int]int)
	dfs = func(i int, m map[int]int) {
		//fmt.Println(m)
		if i == n+1 {
			count++
		}
		for j := 1; j <= n; j++ {
			if m[j] != 0 {
				if i%j == 0 || j%i == 0 {
					m[j]--
					dfs(i+1, m)
					m[j]++
				}
			}
		}
	}
	dfs(1, m)
	return count
}
