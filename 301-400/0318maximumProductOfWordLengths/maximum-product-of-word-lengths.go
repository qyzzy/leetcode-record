package maximumProductOfWordLengths

func maxProduct(words []string) int {
	rec := make([]int, len(words))
	for i, w := range words {
		for _, c := range w {
			rec[i] |= 1 << (c - 'a')
		}
	}

	var ans int
	for i := range rec {
		for j := i + 1; j < len(rec); j++ {
			if rec[i]&rec[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
