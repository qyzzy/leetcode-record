package expressionAddOperators

import "strconv"

func addOperators(num string, target int) []string {
	res := []string{}
	var dfs func(p int, t string, sum int, prevexp int)
	dfs = func(p int, t string, sum int, prevexp int) {
		if p == len(num) {
			//fmt.Println(sum, target, prevexp, p, t)
			if sum == target {
				res = append(res, t)
			}
			return
		}
		for i := p; i < len(num); i++ {
			if num[p] == '0' && p != i {
				break
			}
			n, _ := strconv.Atoi(num[p : i+1])
			if len(t) == 0 {
				dfs(i+1, num[p:i+1], sum+n, n)
			} else {
				dfs(i+1, t+"+"+num[p:i+1], sum+n, n)
				dfs(i+1, t+"-"+num[p:i+1], sum-n, -n)
				dfs(i+1, t+"*"+num[p:i+1], sum-prevexp+prevexp*n, n*prevexp)
			}
		}
	}
	dfs(0, "", 0, 0)
	return res
}
