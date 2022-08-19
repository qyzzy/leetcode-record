package sumRootToLeafNumbers

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
func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, sum int) int {
	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}
	if root.Left == nil {
		return helper(root.Right, sum*10+root.Val)
	}
	if root.Right == nil {
		return helper(root.Left, sum*10+root.Val)
	}
	return helper(root.Right, sum*10+root.Val) + helper(root.Left, sum*10+root.Val)
}
