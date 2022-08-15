package rotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo , something needs to be subtracted

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	length := 0
	x := head
	for x != nil {
		x = x.Next
		length++
	}
	k = k % length
	if k == 0 {
		return head
	}
	dummy := &ListNode{0, head}
	p := dummy
	h := head
	for i := 0; i < k; i++ {
		head = head.Next
	}
	for head.Next != nil {
		head = head.Next
		p = p.Next
	}
	p = p.Next
	dummy.Next = p.Next
	head.Next = h
	p.Next = nil
	return dummy.Next
}
