package licenseKeyFormatting

func licenseKeyFormatting(s string, k int) string {
	res := []byte{}
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			res = append(res, toUpper(s[i]))
			cnt++
			if cnt%k == 0 {
				res = append(res, '-')
			}
		}
	}
	if len(res) > 0 && res[len(res)-1] == '-' {
		res = res[:len(res)-1]
	}
	//fmt.Println(res)
	for i := 0; i < len(res)/2; i++ {
		t := res[i]
		res[i] = res[len(res)-i-1]
		res[len(res)-i-1] = t
	}
	return string(res)
}

func toUpper(b byte) byte {
	if b <= 'z' && b >= 'a' {
		b = byte(int(b - 'a' + 'A'))
	}
	return b
}
