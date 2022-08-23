package summaryRanges

import "strconv"

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	base := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			if i == len(nums)-1 {
				res = append(res, strconv.Itoa(base)+"->"+strconv.Itoa(nums[i]))
			}
			continue
		} else {
			if base != nums[i-1] {
				res = append(res, strconv.Itoa(base)+"->"+strconv.Itoa(nums[i-1]))
			} else {
				res = append(res, strconv.Itoa(base))
			}
			base = nums[i]
			if i == len(nums)-1 {
				res = append(res, strconv.Itoa(base))
			}
		}
	}
	return res
}
