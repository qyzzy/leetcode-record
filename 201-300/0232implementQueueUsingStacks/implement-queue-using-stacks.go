package implementQueueUsingStacks

type MyQueue struct {
	inS, outS []int
}

func Constructor() MyQueue {
	return MyQueue{
		inS:  make([]int, 0),
		outS: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inS = append(this.inS, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outS) == 0 {
		for i := len(this.inS) - 1; i >= 0; i-- {
			this.outS = append(this.outS, this.inS[i])
			this.inS = this.inS[:len(this.inS)-1]
		}
	}
	val := this.outS[len(this.outS)-1]
	this.outS = this.outS[:len(this.outS)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.outS) == 0 {
		for i := len(this.inS) - 1; i >= 0; i-- {
			this.outS = append(this.outS, this.inS[i])
			this.inS = this.inS[:len(this.inS)-1]
		}
	}
	val := this.outS[len(this.outS)-1]
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.inS) == 0 && len(this.outS) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
