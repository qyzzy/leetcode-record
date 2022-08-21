package binarySearchTreeIterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	index int
	data  []int
}

func Constructor(root *TreeNode) BSTIterator {
	data := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		data = append(data, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return BSTIterator{0, data}
}

func (this *BSTIterator) Next() int {
	val := this.data[this.index]
	this.index++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.data)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
