package longestPalindrome

func longestPalindrome(s string) int {
	m := map[byte]int{}
	res := 0
	for i, _ := range s {
		m[s[i]]++
	}
	mid := 0
	for _, v := range m {
		if v == 1 {
			if mid >= 1 {
				continue
			}
			res += v
			mid++
		} else {
			if v%2 == 1 && mid == 0 {
				res += v
				mid++
			} else {
				res += v - v%2
			}
		}
	}
	return res
}
