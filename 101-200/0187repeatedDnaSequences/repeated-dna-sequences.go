package repeatedDnaSequences

func findRepeatedDnaSequences(s string) []string {
	i := 0
	res := []string{}
	m := map[string]int{}
	for i+9 < len(s) {
		tmp := s[i : i+10]
		m[tmp]++
		if m[tmp] == 2 {
			res = append(res, tmp)
		}
		i++
	}
	return res
}
