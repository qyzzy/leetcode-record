package permutationSequence

import "strconv"

func getPermutation(n int, k int) string {
	var num []string
	for i := 1; i <= n; i++ {
		num = append(num, strconv.Itoa(i))
	}
	return permutation(num, k)
}

func permutation(num []string, k int) string {
	if len(num) == 2 {
		if k == 1 {
			return num[0] + num[1]
		} else {
			return num[1] + num[0]
		}
	}
	n := len(num)
	res := ""
	base := 1
	for i := 1; i <= n-1; i++ {
		base *= i
	}
	for i := 0; i < len(num); i++ {
		if k > base {
			k -= base
			continue
		}
		res += num[i]
		num = append(num[:i], num[i+1:]...)
		res += permutation(num, k)
		return res
	}
	return res
}
