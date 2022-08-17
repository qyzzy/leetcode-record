package nQueens

import "strings"

func mak(i int, n int) string {
	res := ""
	for j := 0; j < n; j++ {
		if i == j {
			res += "Q"
		}
		if j == n-1 {
			break
		}
		res += "."
	}
	return res
}

func check(t []string, j int, i int) bool {
	for k := 0; k < len(t); k++ {
		prevj := strings.Index(t[k], "Q")
		if j-prevj == i-k {
			return false
		}
		if i-k == prevj-j {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	var res [][]string
	if n == 1 {
		return [][]string{{"Q"}}
	}
	r := make(map[int]int, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}
	var dfs func(i int, t []string, r map[int]int, prevj int)
	dfs = func(i int, t []string, r map[int]int, prevj int) {
		if i == n {
			tmp := make([]string, n)
			copy(tmp, t)
			res = append(res, tmp)
		}
		s := ""
		for j := 0; j < n; j++ {
			if r[j] == 1 {
				if j == prevj+1 {
					continue
				}
				if !check(t, j, i) {
					continue
				}
				s = mak(j, n)
				r[j] = 0
				t = append(t, s)
				dfs(i+1, t, r, j)
				t = t[:len(t)-1]
				r[j] = 1
			}
		}
	}
	dfs(0, []string{}, r, -100)

	return res
}
