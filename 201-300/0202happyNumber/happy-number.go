package happyNumber

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if m[sum] {
			return false
		}
		m[sum] = true
		n = sum
	}
	return true
}
