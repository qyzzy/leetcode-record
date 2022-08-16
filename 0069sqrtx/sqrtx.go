package sqrtx

func mySqrt(x int) int {
	for i := 1; ; i++ {
		if i*i > x {
			return i - 1
		}
	}
}
