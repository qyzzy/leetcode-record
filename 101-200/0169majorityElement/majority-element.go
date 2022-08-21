package majorityElement

func majorityElement(nums []int) int {
	res := 0
	cnt := 0
	for _, v := range nums {
		if cnt == 0 {
			res = v
			cnt = 1
			continue
		}
		if res != v {
			cnt--
		} else {
			cnt++
		}
	}
	return res
}
