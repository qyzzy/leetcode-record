package convertSortedListToBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	b := make([]*ListNode, 0)
	p := head
	for p != nil {
		b = append(b, p)
		p = p.Next
	}
	root := &TreeNode{b[len(b)/2].Val, nil, nil}
	if len(b)/2 != 0 {
		b[len(b)/2-1].Next = nil
		root.Left = sortedListToBST(b[0])
	}
	if len(b)/2+1 <= len(b)-1 {
		b[len(b)-1].Next = nil
		root.Right = sortedListToBST(b[len(b)/2+1])
	}
	return root
}
