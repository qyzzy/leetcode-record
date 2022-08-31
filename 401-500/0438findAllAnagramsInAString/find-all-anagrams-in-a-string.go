package findAllAnagramsInAString

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	m := map[byte]int{}
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	res := []int{}
	r, l := 0, 0
	for i := 0; i < len(p); i++ {
		m[s[r]]--
		r++
	}
	if judge(m) {
		res = append(res, l)
	}
	m[s[l]]++
	l++
	for r < len(s) {
		m[s[r]]--
		r++
		//fmt.Println(m)
		if judge(m) {
			res = append(res, l)
		}
		m[s[l]]++
		l++
	}
	return res
}

func judge(m map[byte]int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
