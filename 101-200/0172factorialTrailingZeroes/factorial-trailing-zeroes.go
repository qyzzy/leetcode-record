package factorialTrailingZeroes

func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}

//  2 5 10 15 20 25 30
