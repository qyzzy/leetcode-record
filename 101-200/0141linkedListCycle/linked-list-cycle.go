package linkedListCycle

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

//func hasCycle(head *ListNode) bool {
//	m := map[*ListNode]bool{}
//	for head != nil {
//		if _, ok := m[head]; ok {
//			return true
//		}
//		m[head] = true
//		head = head.Next
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p, q := head, head.Next
	for p != q {
		if q.Next == nil || q.Next.Next == nil {
			return false
		}
		p = p.Next
		q = q.Next.Next
	}
	return true
}
