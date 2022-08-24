package rangeSumQueryImmutable

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] += sum[i-1] + nums[i]
	}
	//fmt.Println(sum)
	return NumArray{
		sum: sum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
