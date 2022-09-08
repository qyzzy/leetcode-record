package prefixsum

import (
	"fmt"
	"testing"
)

func Test_twoD(t *testing.T) {
	nums := [][]int{{1, 2, 4, 3}, {5, 1, 2, 4}, {6, 3, 5, 9}}
	twoD := NewTwoD(nums)
	fmt.Println(twoD.ps)
	res := twoD.PrefixSumCompute([]int{1, 1, 2, 2})
	fmt.Println(res)
	res = twoD.PrefixSumCompute([]int{0, 0, 1, 1})
	fmt.Println(res)
	res = twoD.PrefixSumCompute([]int{1, 0, 2, 2})
	fmt.Println(res)
	res = twoD.PrefixSumCompute([]int{0, 1, 2, 2})
	fmt.Println(res)
}
