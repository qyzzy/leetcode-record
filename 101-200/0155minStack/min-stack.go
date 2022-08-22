package minStack

type MinStack struct {
	MinS, S []int
}

func Constructor() MinStack {
	return MinStack{
		MinS: make([]int, 0),
		S:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.MinS) != 0 {
		this.MinS = append(this.MinS, min(val, this.MinS[len(this.MinS)-1]))
	} else {
		this.MinS = append(this.MinS, val)
	}
	this.S = append(this.S, val)
}

func (this *MinStack) Pop() {
	this.MinS = this.MinS[:len(this.MinS)-1]
	this.S = this.S[:len(this.S)-1]
}

func (this *MinStack) Top() int {
	val := this.S[len(this.S)-1]
	return val
}

func (this *MinStack) GetMin() int {
	val := this.MinS[len(this.MinS)-1]
	return val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
