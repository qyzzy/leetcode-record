package numberOfDigitOne

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	if n < 10 && n > 1 {
		return 1
	}
	a := n
	m := 0
	for a > 9 {
		a /= 10
		m++
	}
	count := countDigitOne(n%(pow(m)*a)) + countDigitOne(pow(m)-1)*a
	if a == 1 {
		count += n%(pow(m)*a) + 1
	} else {
		count += pow(m)
	}
	return count
}

func pow(a int) int {
	if a == 0 {
		return 1
	}
	res := 1
	for a > 0 {
		res *= 10
		a--
	}
	return res
}
