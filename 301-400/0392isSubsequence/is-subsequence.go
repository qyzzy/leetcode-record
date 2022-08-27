package isSubsequence

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	sp, tp := 0, 0
	for tp <= len(t)-1 {
		if sp == len(s) {
			return true
		}
		if s[sp] == t[tp] {
			sp++
			tp++
		} else {
			tp++
		}
	}
	if sp == len(s) {
		return true
	}
	return false
}
