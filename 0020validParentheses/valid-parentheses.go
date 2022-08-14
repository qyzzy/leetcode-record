package validParentheses

func isValid(s string) bool {
	i := 0
	count := 0
	var stack [10000]int
	for i < len(s) {
		if s[i] == '(' {
			stack[count] = int('(')
			count++
		}
		if s[i] == '[' {
			stack[count] = int('[')
			count++
		}
		if s[i] == '{' {
			stack[count] = int('{')
			count++
		}
		if s[i] == ')' {
			if count-1 < 0 {
				return false
			}
			if stack[count-1] == int('(') {
				count--
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if count-1 < 0 {
				return false
			}
			if stack[count-1] == int('[') {
				count--
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if count-1 < 0 {
				return false
			}
			if stack[count-1] == int('{') {
				count--
			} else {
				return false
			}
		}
		i++
	}
	return count == 0
}
