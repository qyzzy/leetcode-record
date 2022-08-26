package validPerfectSquare

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	l, r, sum := 2, num/2, 0
	for l <= r {
		sum = (l + (r-l)/2) * (l + (r-l)/2)
		if sum == num {
			return true
		}
		if sum < num {
			l = l + (r-l)/2 + 1
		} else {
			r = l + (r-l)/2 - 1
		}
	}
	return false
}
