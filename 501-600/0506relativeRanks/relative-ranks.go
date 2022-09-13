package relativeRanks

func findRelativeRanks(score []int) []string {
	n := len(score)
	m := make(map[int]int, n)
	for i, v := range score {
		m[v] = i
	}
	sort.Ints(score)
	res := make([]string, n)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			res[m[score[i]]] = "Gold Medal"
			cnt++
		} else if i == n-2 {
			res[m[score[i]]] = "Silver Medal"
			cnt++
		} else if i == n-3 {
			res[m[score[i]]] = "Bronze Medal"
			cnt++
		} else {
			cnt++
			res[m[score[i]]] = strconv.Itoa(cnt)
		}
	}
	return res
}
