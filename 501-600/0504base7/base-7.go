package base7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	res := []byte{}
	for num > 0 {
		res = append(res, byte(num%7+'0'))
		num /= 7
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	if flag {
		return "-" + string(res)
	}
	return string(res)
}
