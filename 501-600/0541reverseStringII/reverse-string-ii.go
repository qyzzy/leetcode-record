package reverseStringII

func reverse(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		tmp := b[l]
		b[l] = b[r]
		b[r] = tmp
		l++
		r--
	}
	return
}

func reverseStr(s string, k int) string {
	res := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k < len(s) {
			reverse(res[i : i+k])
		} else {
			reverse(res[i:len(s)])
		}
	}
	return string(res)
}
