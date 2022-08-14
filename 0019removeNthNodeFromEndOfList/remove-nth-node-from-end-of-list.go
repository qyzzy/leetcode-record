package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p := dummy
	for i := 0; i < n; i++ {
		head = head.Next
	}
	for head != nil {
		p = p.Next
		head = head.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
