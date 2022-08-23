package palindromeLinkedList

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
func isPalindrome(head *ListNode) bool {
	b := make([]int, 0)
	for head != nil {
		b = append(b, head.Val)
		head = head.Next
	}
	return isP(b)
}

func isP(a []int) bool {
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}
