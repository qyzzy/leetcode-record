package wordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	tmp := strings.Split(s, " ")
	if len(tmp) != len(pattern) {
		return false
	}
	m := map[byte]string{}
	n := map[string]byte{}
	for i, v := range tmp {
		if val, ok := m[pattern[i]]; ok {
			if val != v {
				return false
			}
		} else {
			m[pattern[i]] = v
		}
		if val, ok := n[v]; ok {
			if pattern[i] != val {
				return false
			}
		} else {
			n[v] = pattern[i]
		}
	}
	return true
}
