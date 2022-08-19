package gasStation

func canCompleteCircuit(gas []int, cost []int) int {
	res := 0
	cur, sum := 0, 0
	for i := 0; i < len(gas); i++ {
		cur += gas[i] - cost[i]
		sum += gas[i] - cost[i]
		if cur < 0 {
			cur = 0
			res = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	return res
}
