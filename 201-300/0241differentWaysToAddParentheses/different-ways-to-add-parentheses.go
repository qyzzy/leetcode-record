package differentWaysToAddParentheses

import "strconv"

func diffWaysToCompute(expression string) []int {
	if isD(expression) {
		res, _ := strconv.Atoi(expression)
		return []int{res}
	}
	res := []int{}
	for i, v := range expression {
		if v == '+' || v == '-' || v == '*' {
			l, r := diffWaysToCompute(expression[:i]), diffWaysToCompute(expression[i+1:])
			for _, vl := range l {
				for _, vr := range r {
					sum := 0
					if v == '+' {
						sum = vl + vr
					}
					if v == '-' {
						sum = vl - vr
					}
					if v == '*' {
						sum = vl * vr
					}
					res = append(res, sum)
				}
			}
		}
	}
	return res
}

func isD(e string) bool {
	for i := 0; i < len(e); i++ {
		if e[i] >= '0' && e[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
