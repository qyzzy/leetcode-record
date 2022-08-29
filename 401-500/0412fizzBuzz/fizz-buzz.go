package fizzBuzz

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	tmp := ""
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		res[i-1] = tmp
		tmp = ""
	}
	return res
}
