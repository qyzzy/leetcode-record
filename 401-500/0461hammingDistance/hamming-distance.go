package hammingDistance

func hammingDistance(x int, y int) int {
	sum, res := x^y, 0
	for sum > 0 {
		res++
		sum &= sum - 1
	}
	return res
}
