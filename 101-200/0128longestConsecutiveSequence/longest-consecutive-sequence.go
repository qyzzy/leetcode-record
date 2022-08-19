package longestConsecutiveSequence

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	res := 0
	for k, _ := range m {
		if m[k-1] {
			continue
		}
		tmp := k
		for m[tmp+1] {
			tmp++
		}
		res = max(res, tmp-k+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
