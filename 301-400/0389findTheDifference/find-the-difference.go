package findTheDifference

func findTheDifference(s string, t string) byte {
	res := 0
	for i := 0; i < len(s); i++ {
		res += int(t[i] - s[i])
	}
	return byte(res + int(t[len(t)-1]))
}
