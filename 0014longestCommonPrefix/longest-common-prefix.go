package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]
	for _, v := range strs {
		pref = prefix(v, pref)
		//fmt.Println(pref)
	}
	return pref
}

func prefix(s1, s2 string) string {
	if s1 == "" || s2 == "" {
		return ""
	}
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			if i == 0 {
				return ""
			} else {
				return s1[:i]
			}
		}
	}
	if len(s1) < len(s2) {
		return s1
	}
	return s2
}
