package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	mid := middle(head)
	r := mid.Next
	mid.Next = nil
	r = reverse(r)
	head = merge(head, r)
}

func middle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

func merge(a, b *ListNode) *ListNode {
	res := &ListNode{0, a}
	p := res.Next
	a = a.Next
	for a != nil && b != nil {
		p.Next = b
		b = b.Next
		p = p.Next
		p.Next = a
		a = a.Next
		p = p.Next
	}
	if a != nil {
		p.Next = a
	}
	if b != nil {
		p.Next = b
	}
	return res.Next
}

func reverse(head *ListNode) *ListNode {
	var p *ListNode
	c := head
	for c != nil {
		n := c.Next
		c.Next = p
		p = c
		c = n
	}
	return p
}
