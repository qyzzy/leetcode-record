package assignCookies

import "sort"

func findContentChildren(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	for i := 0; i < len(g) && j < len(s); i++ {
		for j < len(s) && g[i] > s[j] {
			j++
		}
		if j < len(s) {
			if g[i] <= s[j] {
				res++
				j++
			}
		}
	}
	return res
}
