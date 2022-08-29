package reconstructOriginalDigitsFromEnglish

import "sort"

func originalDigits(s string) string {
	// g -> 8
	// x -> 6
	// z -> 0
	// w -> 2
	// h - 8 -> 3
	// u -> 4
	// f - 4 -> 5
	// s - 6 -> 7
	// i - 5 - 6 - 8 -> 9
	// n - 7 - 2*9 -> 1
	cnt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	res := []byte{}
	for cnt['g'] > 0 {
		cnt['e']--
		cnt['i']--
		cnt['g']--
		cnt['h']--
		cnt['t']--
		res = append(res, '8')
	}
	for cnt['x'] > 0 {
		cnt['s']--
		cnt['i']--
		cnt['x']--
		res = append(res, '6')
	}
	for cnt['z'] > 0 {
		cnt['z']--
		cnt['e']--
		cnt['r']--
		cnt['o']--
		res = append(res, '0')
	}
	for cnt['w'] > 0 {
		cnt['t']--
		cnt['w']--
		cnt['o']--
		res = append(res, '2')
	}
	for cnt['u'] > 0 {
		cnt['f']--
		cnt['o']--
		cnt['u']--
		cnt['r']--
		res = append(res, '4')
	}
	for cnt['h'] > 0 {
		cnt['t']--
		cnt['h']--
		cnt['r']--
		cnt['e'] -= 2
		res = append(res, '3')
	}
	for cnt['f'] > 0 {
		cnt['f']--
		cnt['i']--
		cnt['v']--
		cnt['e']--
		res = append(res, '5')
	}
	for cnt['s'] > 0 {
		cnt['s']--
		cnt['e'] -= 2
		cnt['v']--
		cnt['n']--
		res = append(res, '7')
	}
	for cnt['i'] > 0 {
		cnt['n'] -= 2
		cnt['i']--
		cnt['e']--
		res = append(res, '9')
	}
	for cnt['n'] > 0 {
		cnt['o']--
		cnt['n']--
		cnt['e']--
		res = append(res, '1')
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return string(res)
}
