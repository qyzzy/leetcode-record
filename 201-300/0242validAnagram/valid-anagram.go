package validAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[byte]int{}
	for _, v := range s {
		m[byte(v)]++
	}
	for _, v := range t {
		if _, ok := m[byte(v)]; ok {
			m[byte(v)]--
			if m[byte(v)] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
