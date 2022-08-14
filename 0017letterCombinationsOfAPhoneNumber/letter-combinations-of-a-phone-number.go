package letterCombinationsOfAPhoneNumber

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var res []string
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(p int, t string)
	dfs = func(p int, t string) {
		if p >= len(digits) {
			res = append(res, t)
			return
		}
		for i := 0; i < len(m[digits[p]]); i++ {
			dfs(p+1, t+m[digits[p]][i:i+1])
		}
	}
	dfs(0, "")
	return res
}
