package prefixsum

type PrefixSum interface {
	PrefixSumCompute(indexes []int) int
}

type TwoD struct {
	ps [][]int
}

func NewTwoD(nums [][]int) *TwoD {
	for i := 1; i < len(nums); i++ {
		nums[i][0] = nums[i][0] + nums[i-1][0]
	}

	for i := 1; i < len(nums[0]); i++ {
		nums[0][i] = nums[0][i] + nums[0][i-1]
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < len(nums[i]); j++ {
			nums[i][j] = nums[i-1][j] + nums[i][j-1] - nums[i-1][j-1] + nums[i][j]
		}
	}
	return &TwoD{
		ps: nums,
	}
}

func (td *TwoD) PrefixSumCompute(indexes []int) int {
	x1, y1, x2, y2 := indexes[0], indexes[1], indexes[2], indexes[3]
	if x1 == 0 && y1 == 0 {
		return td.ps[x2][y2]
	}
	if x1 == 0 {
		return td.ps[x2][y2] - td.ps[x2][0]
	}
	if y1 == 0 {
		return td.ps[x2][y2] - td.ps[0][y2]
	}
	return td.ps[x2][y2] - td.ps[x2][y1-1] - td.ps[x1-1][y2] + td.ps[x1-1][y1-1]
}
