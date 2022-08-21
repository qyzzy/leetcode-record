package bitwiseAndOfNumbersRange

// Find the longest common prefix

func rangeBitwiseAnd(left int, right int) int {
	cnt := 0
	for left < right {
		left = left >> 1
		right = right >> 1
		cnt++
	}
	return left << cnt
}
