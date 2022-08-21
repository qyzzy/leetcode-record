package removeLinkedListElements

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
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	n := removeElements(head.Next, val)
	if head.Val == val {
		return n
	} else {
		head.Next = n
	}
	return head
}
