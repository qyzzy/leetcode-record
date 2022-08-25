package rangeSumQueryMutable

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	return NumArray{
		sum: sum,
	}
}

func (this *NumArray) Update(index int, val int) {
	dif := val - this.sum[index+1] + this.sum[index]
	for i := index + 1; i < len(this.sum); i++ {
		this.sum[i] += dif
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
