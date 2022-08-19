package validPalindrome

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !judge(s[l]) {
			l++
		}
		for l < r && !judge(s[r]) {
			r--
		}
		if convert(s[l]) != convert(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func convert(b byte) byte {
	if b >= 'a' && b <= 'z' {
		b = byte(b - 'a' + 'A')
	}
	return b
}

func judge(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}
