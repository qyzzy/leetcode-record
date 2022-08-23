package uglyNumber

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	f := []int{2, 3, 5}
	for i := 0; i < 3; i++ {
		for n%f[i] == 0 {
			n /= f[i]
		}
	}
	return n == 1
}
