package restoreIpAddresses

import "strconv"

func restoreIpAddresses(s string) []string {
	var res []string
	dfs(s, "", 0, 0, &res)
	return res
}

func dfs(s string, t string, prev int, para int, res *[]string) {
	if prev == len(s) && para == 4 {
		*res = append(*res, t[1:])
	}
	if para == 4 {
		return
	}
	for i := prev; i < len(s) && i-prev <= 2; i++ {
		tmp := s[prev : i+1]
		num, _ := strconv.Atoi(string(tmp))
		if len(tmp) > 1 {
			if tmp[0] == '0' {
				continue
			}
		}
		if num >= 0 && num <= 255 {
			t += "." + s[prev:i+1]
		} else {
			continue
		}
		dfs(s, t, i+1, para+1, res)
		t = t[:len(t)-(i-prev)-2]
	}
}
