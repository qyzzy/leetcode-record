package stringCompression

func compress(chars []byte) int {
	pre := chars[0]
	ind := 0
	cnt := 1
	for i := 1; i < len(chars); {
		if chars[i] == pre {
			cnt++
			if cnt == 2 {
				ind = i
			}
			if judge(cnt) {
				i++
				continue
			}
			if cnt > 2 {
				chars = append(chars[:i], chars[i+1:]...)
			} else {
				i++
			}
		} else {
			if cnt > 1 {
				tmp := []byte(strconv.Itoa(cnt))
				for j := 0; j < len(tmp); j++ {
					chars[ind+j] = tmp[j]
				}
				cnt = 1
			}
			pre = chars[i]
			i++
		}
	}
	if cnt > 1 {
		tmp := []byte(strconv.Itoa(cnt))
		for j := 0; j < len(tmp); j++ {
			chars[ind+j] = tmp[j]
		}
		cnt = 1
	}
	return len(chars)
}

func judge(num int) bool {
	if num == 10 || num == 100 || num == 1000 {
		return true
	}
	return false
}
