package isomorphicStrings

func isIsomorphic(s string, t string) bool {
	st := make([]int, 256)
	for i := 0; i < len(st); i++ {
		st[i] = 257
	}
	ts := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if st[int(s[i])] == 257 {
			st[int(s[i])] = int(s[i]) - int(t[i])
		}
		if st[int(s[i])] != int(s[i])-int(t[i]) {
			return false
		}
		if ts[int(t[i])] == 0 {
			ts[int(t[i])] = int(s[i])
		}
		if ts[int(t[i])] != int(s[i]) {
			return false
		}
	}
	return true
}
