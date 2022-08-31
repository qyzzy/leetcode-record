package arrangingCoins

func arrangeCoins(n int) int {
	c := n
	for i := 1; i <= n; i++ {
		if c >= i {
			c -= i
		} else {
			return i - 1
		}
	}
	return 1
}
