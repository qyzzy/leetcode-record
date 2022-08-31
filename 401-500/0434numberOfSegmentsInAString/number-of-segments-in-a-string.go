package numberOfSegmentsInAString

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != ' ' && s[i+1] == ' ' {
			res++
		}
	}
	if s[len(s)-1] != ' ' {
		res++
	}
	return res

}
