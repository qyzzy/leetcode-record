package slidingWindowMaximum

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

/*
	queue
*/
//func maxSlidingWindow(nums []int, k int) []int {
//	res := make([]int, len(nums) - k + 1)
//	st := []int{}
//	for i := 0; i < len(nums); i++ {
//		if len(st) == 0 {
//			st = append(st, i)
//		}   else {
//			st = append(st, i)
//			j := len(st) - 2
//			for len(st) > 0 && j >= 0 && nums[i] > nums[st[j]] {
//				st = append(st[:j], st[j + 1:]...)
//				j--
//			}
//			for len(st) > 0 && st[0] < i - k + 1 {
//				st = st[1:]
//			}
//		}
//		if i >= k - 1 {
//			res[i - k + 1] = nums[st[0]]
//		}
//		//fmt.Println(st, res)
//	}
//	return res
//}
