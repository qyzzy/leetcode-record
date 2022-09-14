package detectCapital

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	first := 0
	if word[0] <= 'Z' && word[0] >= 'A' {
		first = 1
	}
	flag := 0
	if word[1] <= 'Z' && word[1] >= 'A' {
		flag = 1
	}
	if first == 0 && flag == 1 {
		return false
	}
	for i := 2; i < len(word); i++ {
		if (flag == 1 && word[i] <= 'Z' && word[i] >= 'A') || (flag == 0 && word[i] <= 'z' && word[i] >= 'a') {
			continue
		} else {
			return false
		}
	}
	return true
}
