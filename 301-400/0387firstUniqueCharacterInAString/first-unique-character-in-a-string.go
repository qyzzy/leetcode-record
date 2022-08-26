package firstUniqueCharacterInAString

func firstUniqChar(s string) int {
	b := make([]int, 27)
	for i := 0; i < len(s); i++ {
		b[int(s[i]-'a')]++
	}
	for i := 0; i < len(s); i++ {
		if b[int(s[i]-'a')] == 1 {
			return i
		}
	}
	return -1
}
