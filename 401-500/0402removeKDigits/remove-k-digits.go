package removeKDigits

func removeKdigits(num string, k int) string {
	st := []byte{}
	for i := 0; i < len(num); i++ {
		if len(st) == 0 {
			st = append(st, num[i])
		} else {
			for len(st) > 0 && st[len(st)-1] > num[i] && k > 0 {
				st = st[:len(st)-1]
				k--
			}
			st = append(st, num[i])
		}
	}
	for k > 0 {
		st = st[:len(st)-1]
		k--
	}
	for i := 0; i < len(st); {
		if st[0] == '0' && i != len(st)-1 {
			st = st[1:]
		} else {
			i++
		}
	}
	if len(st) == 0 {
		return "0"
	}
	return string(st)
}
