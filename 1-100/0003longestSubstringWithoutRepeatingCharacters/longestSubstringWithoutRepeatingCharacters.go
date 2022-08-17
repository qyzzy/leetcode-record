package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := map[byte]int{}
	l, r := 0, 0
	for r < len(s) {
		m[s[r]]++
		for m[s[r]] >= 2 {
			m[s[l]]--
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
