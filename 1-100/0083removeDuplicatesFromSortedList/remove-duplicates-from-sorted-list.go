package removeDuplicatesFromSortedList

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
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		q := p.Next
		if q != nil && p.Val == q.Val {
			p.Next = q.Next
		} else {
			p = p.Next
		}
	}
	return head
}
