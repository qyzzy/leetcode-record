package fibonacciNumber

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	p, q := 0, 1
	for i := 2; i <= n; i++ {
		t := q
		q = q + p
		p = t
	}
	return q
}
