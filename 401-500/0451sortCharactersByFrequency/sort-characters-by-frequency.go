package sortCharactersByFrequency

import "sort"

func frequencySort(s string) string {
	m := [][]int{}
	for i := 0; i < 150; i++ {
		t := make([]int, 2)
		m = append(m, t)
	}
	for i := 0; i < len(s); i++ {
		m[int(s[i])][0] = int(s[i])
		m[int(s[i])][1]++
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i][1] > m[j][1]
	})
	res := []byte{}
	for i := 0; i < 150; i++ {
		if m[i][1] > 0 {
			for m[i][1] > 0 {
				m[i][1]--
				res = append(res, byte(m[i][0]))
			}
		} else {
			break
		}
	}
	return string(res)
}
