package substringWithConcatenationOfAllWords

func findSubstring(s string, words []string) []int {
	var res []int
	n := len(words[0])
	m := map[string]int{}
	for _, v := range words {
		m[v]++
	}
	l, r := -1, len(words)*n-2
	for r < len(s)-1 {
		l++
		r++
		if check(m, s[l:r+1], n) {
			res = append(res, l)
		}
	}
	return res
}

func check(m map[string]int, s string, n int) bool {
	sw := map[string]int{}
	for i := 0; i < len(s); i += n {
		sw[s[i:i+n]]++
	}
	for k, v := range m {
		if v1, ok := sw[k]; ok {
			if v1 != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
