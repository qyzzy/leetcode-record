package reverseString

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		t := s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = t
	}
}
