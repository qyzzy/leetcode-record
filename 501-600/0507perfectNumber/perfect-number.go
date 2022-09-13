package perfectNumber

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := -num
	for i := 1; ; i++ {
		if i*i > num {
			break
		}
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
}
