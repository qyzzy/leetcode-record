package minimumWindowSubstring

func minWindow(s string, t string) string {
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	res := ""
	for i := 0; i < len(t); i++ {
		m1[t[i]]++
	}
	r, l := 0, 0
	for r < len(s) {
		m2[s[r]]++
		for check(m1, m2) {
			if res == "" || len(res) > r-l+1 {
				res = s[l : r+1]
			}
			m2[s[l]]--
			l++
		}
		r++
	}
	return res
}

func check(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			if v2 < v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
