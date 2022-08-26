package shuffleAnArray

import "math/rand"

type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int(nil), nums...)}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}

func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	for i := range s.nums {
		j := i + rand.Intn(n-i)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
