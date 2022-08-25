package oddEvenLinkedList

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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, pp := head, head.Next
	left := head.Next
	for pp != nil && pp.Next != nil {
		p.Next = pp.Next
		p = p.Next
		pp.Next = p.Next
		pp = pp.Next
	}
	p.Next = left
	return head
}
