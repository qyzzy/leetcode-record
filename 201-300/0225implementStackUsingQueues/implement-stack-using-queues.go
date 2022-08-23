package implementStackUsingQueues

type MyStack struct {
	inQ, outQ []int
}

func Constructor() MyStack {
	return MyStack{
		inQ:  make([]int, 0),
		outQ: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.inQ = append(this.inQ, x)
}

func (this *MyStack) Pop() int {
	if len(this.outQ) == 0 {
		for i := 0; i < len(this.inQ); {
			this.outQ = append(this.outQ, this.inQ[i])
			this.inQ = this.inQ[1:]
		}
	}
	val := this.outQ[len(this.outQ)-1]
	this.outQ = this.outQ[:len(this.outQ)-1]
	return val
}

func (this *MyStack) Top() int {
	if len(this.inQ) != 0 {
		for i := 0; i < len(this.inQ); {
			this.outQ = append(this.outQ, this.inQ[i])
			this.inQ = this.inQ[1:]
		}
	}
	val := this.outQ[len(this.outQ)-1]
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.inQ) == 0 && len(this.outQ) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
