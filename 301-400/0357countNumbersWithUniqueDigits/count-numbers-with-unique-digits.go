package countNumbersWithUniqueDigits

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	sum := 0
	factor := 1
	for i := 1; i < n; i++ {
		factor *= 10 - i
		sum += 9 * factor
	}
	return sum + 10
}
