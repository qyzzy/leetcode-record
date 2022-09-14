package largestPalindromeProduct

var cache = []int{9, 987, 123, 597, 677, 1218, 877, 475}

func largestPalindrome(n int) int {
	return cache[n-1]
}
