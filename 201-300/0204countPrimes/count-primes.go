package countPrimes

func countPrimes(n int) int {
	res := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
