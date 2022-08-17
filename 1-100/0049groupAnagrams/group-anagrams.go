package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	var res [][]string
	for _, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		m[string(tmp)] = append(m[string(tmp)], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
