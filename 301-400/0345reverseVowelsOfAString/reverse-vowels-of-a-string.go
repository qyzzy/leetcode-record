package reverseVowelsOfAString

func reverseVowels(s string) string {
	n := len(s)
	b := []byte(s)
	for l, r := 0, n-1; l < r; {
		for !IsVowel(s[r]) && l < r {
			r--
		}
		for !IsVowel(s[l]) && l < r {
			l++
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func IsVowel(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}
	if s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
